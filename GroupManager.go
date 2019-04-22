package main

import (
	"./Messages"
	"database/sql"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"log"
	"sync"
)

type Group struct {
	clients map[*Client]bool
}

var activeGroups = map[string]*Group{}
var activeGroupsMutex = sync.RWMutex{}

func addUserToGroup(groupName string, username string, sql *sql.DB) error {
	stmt, prepErr := sql.Prepare("INSERT INTO GroupMembers VALUES (?, ?)")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	_, execErr := stmt.Exec(groupName, username)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return execErr
	}
	return nil
}

func createGroup(groupName string, sql *sql.DB) error {
	stmt, prepErr := sql.Prepare("INSERT INTO Groups VALUES (?)")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	_, execErr := stmt.Exec(groupName)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return execErr
	}
	return nil
}

func createGroupHandler(message *Message, sql *sql.DB) {
	var createGroupRunner = func() {
		createGroupMsg := Messages.CreateGroupReq{}
		parseErr := proto.Unmarshal(message.body, &createGroupMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		err := createGroup(createGroupMsg.GroupName, sql)
		if err == nil {
			err = addUserToGroup(createGroupMsg.GroupName, *message.client.username, sql)
			if err == nil {
				err = joinActiveGroup(createGroupMsg.GroupName, message.client, sql)
			}
		}
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("createGroupError", errorData)
			return
		}
	}
	go createGroupRunner()
}

func inviteToGroup(toUsername string, groupName string, fromUsername string, sql *sql.DB) error {
	inviteUUID, uuidErr := uuid.NewRandom()
	if uuidErr != nil {
		log.Fatalln("UUID ERROR: ", uuidErr)
		return uuidErr
	}
	inviteID := inviteUUID.String()
	stmt, prepErr := sql.Prepare("INSERT INTO Invites VALUES (?, ?, ?, ?)")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	_, execErr := stmt.Exec(inviteID, groupName, toUsername, fromUsername)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return execErr
	}
	return nil
}

func inviteToGroupHandler(message *Message, sql *sql.DB) {
	var inviteToGroupRunner = func() {
		inviteToGroupMsg := Messages.InviteReq{}
		parseErr := proto.Unmarshal(message.body, &inviteToGroupMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		err := inviteToGroup(inviteToGroupMsg.Username, *message.client.groupName, *message.client.username, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("inviteError", errorData)
			return
		}
	}
	go inviteToGroupRunner()
}

func getGroupNameFromInviteID(inviteID string, sql *sql.DB) (string, error) {
	stmt, prepErr := sql.Prepare("SELECT groupName FROM Invites WHERE id=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return "", prepErr
	}
	rows, execErr := stmt.Query(inviteID)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return "", execErr
	}
	if rows.Next() {
		var groupName string
		rows.Scan(&groupName)
		return groupName, nil
	} else {
		return "", errors.New("InviteID does not exist")
	}
}

func deleteInvite(inviteID string, sql *sql.DB) error {
	stmt, prepErr := sql.Prepare("DELETE FROM Invites WHERE id=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return prepErr
	}
	_, execErr := stmt.Exec(inviteID)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return execErr
	}
	return nil
}

func acceptInviteHandler(message *Message, sql *sql.DB) {
	var acceptInviteRunner = func() {
		acceptInviteMsg := Messages.AcceptInviteReq{}
		parseErr := proto.Unmarshal(message.body, &acceptInviteMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		groupName, err := getGroupNameFromInviteID(acceptInviteMsg.InviteID, sql)
		if err == nil {
			err = addUserToGroup(groupName, *message.client.username, sql)
			if err == nil {
				err = deleteInvite(acceptInviteMsg.InviteID, sql)
			}
		}
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("acceptInviteError", errorData)
			return
		}

	}
	go acceptInviteRunner()
}

func deleteInviteHandler(message *Message, sql *sql.DB) {
	var deleteInviteRunner = func() {
		deleteInviteMsg := Messages.DeleteInviteReq{}
		parseErr := proto.Unmarshal(message.body, &deleteInviteMsg)
		if parseErr != nil {
			log.Fatalln("PARSE ERROR: ", parseErr)
			return
		}
		err := deleteInvite(deleteInviteMsg.InviteID, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("deleteInviteError", errorData)
			return
		}
	}
	go deleteInviteRunner()
}

func getUserGroups(username string, sql *sql.DB) ([]string, error) {
	stmt, prepErr := sql.Prepare("SELECT groupName FROM GroupMembers WHERE username=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return nil, prepErr
	}
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
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

func getGroupsHandler(message *Message, sql *sql.DB) {
	var getGroupsRunner = func() {
		groupNames, err := getUserGroups(*message.client.username, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("getGroupsError", errorData)
			return
		}
		groupsMsg := Messages.GroupsResp{GroupNames: groupNames}
		groupsData, serializeErr := proto.Marshal(&groupsMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.send("getGroups", groupsData)
	}
	go getGroupsRunner()
}

type Invite struct {
	inviteID string
	fromUsername string
	groupName string
}

func getUserInvites(username string, sql *sql.DB) ([]Invite, error) {
	stmt, prepErr := sql.Prepare("SELECT id, fromUsername, groupName FROM Invites WHERE toUsername=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return nil, prepErr
	}
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return nil, execErr
	}
	var invites []Invite
	for rows.Next() {
		var inviteID string
		var fromUsername string
		var groupName string
		rows.Scan(&inviteID, &fromUsername, &groupName)
		invite := Invite{
			inviteID: inviteID,
			fromUsername: fromUsername,
			groupName: groupName,
		}
		invites = append(invites, invite)
	}
	return invites, nil
}

func getInvitesHandler(message *Message, sql *sql.DB) {
	var getInvitesRunner = func() {
		invites, err := getUserInvites(*message.client.username, sql)
		if err != nil {
			errorMsg := Messages.Error{Message: err.Error()}
			errorData, serializeErr := proto.Marshal(&errorMsg)
			if serializeErr != nil {
				log.Fatalln("SERIALIZE ERROR: ", serializeErr)
				return
			}
			message.client.send("getInvitesError", errorData)
			return
		}

		var protoInvites []*Messages.InvitesResp_Invite
		for _, invite := range invites {
			protoInvite := Messages.InvitesResp_Invite{
				InviteID: invite.inviteID,
				FromUsername: invite.fromUsername,
				GroupName: invite.groupName,
			}
			protoInvites = append(protoInvites, &protoInvite)
		}
		invitesMsg := Messages.InvitesResp{Invites: protoInvites}
		invitesData, serializeErr := proto.Marshal(&invitesMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.send("getInvites", invitesData)
	}
	go getInvitesRunner()
}

func isInGroup(username string, groupName string, sql *sql.DB) (bool, error) {
	stmt, prepErr := sql.Prepare("SELECT groupName FROM GroupMembers WHERE username=?")
	if prepErr != nil {
		log.Fatalln("MySQL Prepare Error: ", prepErr)
		return false, prepErr
	}
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		log.Println("Query Error: ", execErr)
		return false, execErr
	}
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
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
		activeGroup = &Group{}
		activeGroups[groupName] = activeGroup
	}
	activeGroup.clients[client] = true
	*client.groupName = groupName
	activeGroupsMutex.Unlock()
	return nil
}

func joinGroupHandler(message *Message, sql *sql.DB) {
	joinGroupMsg := Messages.JoinGroupReq{}
	parseErr := proto.Unmarshal(message.body, &joinGroupMsg)
	if parseErr != nil {
		log.Fatalln("PARSE ERROR: ", parseErr)
		return
	}
	err := joinActiveGroup(joinGroupMsg.GroupName, message.client, sql)
	if err != nil {
		errorMsg := Messages.Error{Message: err.Error()}
		errorData, serializeErr := proto.Marshal(&errorMsg)
		if serializeErr != nil {
			log.Fatalln("SERIALIZE ERROR: ", serializeErr)
			return
		}
		message.client.send("joinGroupError", errorData)
		return
	}
}