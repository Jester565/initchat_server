package main

import (
	"./Messages"
	"database/sql"
	"errors"
	"github.com/golang/protobuf/proto"
	"sync"
)

type Group struct {
	clients map[*Client]bool
}

var activeGroups = make(map[string]*Group)
var activeGroupsMutex = sync.RWMutex{}

func addGroupToDatabase(groupName string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("INSERT INTO Groups VALUES (?)")
	_, execErr := stmt.Exec(groupName)
	if execErr != nil {
		return execErr
	}
	return nil
}

func joinActiveGroup(groupName string, client *Client, sql *sql.DB) error {
	inGroup, err := isInGroup(*client.username, groupName, sql)
	if err != nil {
		return err
	}
	if !inGroup {
		return errors.New("User is not in group")
	}
	activeGroupsMutex.Lock()
	activeGroup, hasGroup := activeGroups[groupName]
	if !hasGroup {
		activeGroup = &Group{
			clients: make(map [*Client]bool),
		}
		activeGroups[groupName] = activeGroup
	}
	activeGroup.clients[client] = true
	client.groupName = &groupName
	activeGroupsMutex.Unlock()
	addGroupMessage(*client.username + " joined the group", groupName, *client.username, sql)
	return nil
}

func addGroupMemberToDatabase(groupName string, username string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("INSERT INTO GroupMembers VALUES (?, ?)")
	_, execErr := stmt.Exec(groupName, username)
	if execErr != nil {
		return execErr
	}
	return nil
}

func sendGroup(groupName string, client *Client, sql *sql.DB) error {
	messages, err := getGroupMessages(groupName, sql)
	if err != nil {
		return err
	}
	groupMsg := Messages.GroupResp{
		Messages: messages,
	}
	groupData, _ := proto.Marshal(&groupMsg)
	client.send("group", groupData)
	return nil
}

//Adds group to database and the user to the activeGroups map
func CreateGroupHandler(message *Message, sql *sql.DB) {
	go func() {
		createGroupMsg := Messages.CreateGroupReq{}
		proto.Unmarshal(message.body, &createGroupMsg)
		err := addGroupToDatabase(createGroupMsg.GroupName, sql)
		if err == nil {
			err = addGroupMemberToDatabase(createGroupMsg.GroupName, *message.client.username, sql)
			if err == nil {
				err = sendGroup(createGroupMsg.GroupName, message.client, sql)
				if err == nil {
					err = joinActiveGroup(createGroupMsg.GroupName, message.client, sql)
				}
			}
		}
		if err != nil {
			sendError("createGroupErr", err, message.client)
		}
	}()
}

func getUserGroups(username string, sql *sql.DB) ([]string, error) {
	stmt, _ := sql.Prepare("SELECT groupName FROM GroupMembers WHERE username=?")
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		return nil, execErr
	}
	var groupNames []string
	for rows.Next() {
		var groupName string
		rows.Scan(&groupName)
		groupNames = append(groupNames, groupName)
	}
	return groupNames, nil
}

//Get groups the user is a member of
func GetGroupsHandler(message *Message, sql *sql.DB) {
	go func() {
		groupNames, err := getUserGroups(*message.client.username, sql)
		if err != nil {
			sendError("getGroupsErr", err, message.client)
			return
		}
		groupsMsg := Messages.GroupsResp{GroupNames: groupNames}
		groupsData, _ := proto.Marshal(&groupsMsg)
		message.client.send("getGroups", groupsData)
	}()
}

func isInGroup(username string, groupName string, sql *sql.DB) (bool, error) {
	stmt, _ := sql.Prepare("SELECT groupName FROM GroupMembers WHERE username=?")
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		return false, execErr
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

//Add user to activeGroups map
func JoinGroupHandler(message *Message, sql *sql.DB) {
	joinGroupMsg := Messages.JoinGroupReq{}
	proto.Unmarshal(message.body, &joinGroupMsg)
	err := joinActiveGroup(joinGroupMsg.GroupName, message.client, sql)
	if err == nil {
		err = sendGroup(joinGroupMsg.GroupName, message.client, sql)
	}
	if err != nil {
		sendError("joinGroupErr", err, message.client)
		return
	}
}

func leaveActiveGroup(client *Client, sql *sql.DB) {
	//Delete from groups
	if client.groupName != nil {
		addGroupMessage(*client.username + " has left the group", *client.groupName, *client.username, sql)
		activeGroupsMutex.Lock()
		group, hasGroup := activeGroups[*client.groupName]
		if hasGroup {
			delete(group.clients, client)
			if len(group.clients) == 0 {
				delete(activeGroups, *client.groupName)
			}
		}
		activeGroupsMutex.Unlock()
		client.groupName = nil
	}
}

func LeaveGroupHandler(message *Message, sql *sql.DB) {
	leaveActiveGroup(message.client, sql)
}

func RefreshGroupHandler(message *Message, sql *sql.DB) {
	sendGroup(*message.client.groupName, message.client, sql)
}