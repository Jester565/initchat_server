package main

import (
	"./Messages"
	"database/sql"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
)

func addInviteToDatabase(toUsername string, groupName string, fromUsername string, sql *sql.DB) error {
	inviteUUID, _ := uuid.NewRandom()

	inviteID := inviteUUID.String()
	stmt, _ := sql.Prepare("INSERT INTO Invites VALUES (?, ?, ?, ?)")
	_, execErr := stmt.Exec(inviteID, groupName, toUsername, fromUsername)
	if execErr != nil {
		return execErr
	}
	return nil
}

//Creates database entry for the invite
func InviteToGroupHandler(message *Message, sql *sql.DB) {
	go func() {
		inviteToGroupMsg := Messages.InviteReq{}
		proto.Unmarshal(message.body, &inviteToGroupMsg)
		err := addInviteToDatabase(inviteToGroupMsg.Username, *message.client.groupName, *message.client.username, sql)
		if err != nil {
			sendError("inviteErr", err, message.client)
		}
	}()
}

func getGroupNameFromInviteID(inviteID string, sql *sql.DB) (string, error) {
	stmt, _ := sql.Prepare("SELECT groupName FROM Invites WHERE id=?")
	rows, execErr := stmt.Query(inviteID)
	if execErr != nil {
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

func deleteInviteFromDatabase(inviteID string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("DELETE FROM Invites WHERE id=?")
	_, execErr := stmt.Exec(inviteID)
	if execErr != nil {
		return execErr
	}
	return nil
}

func AcceptInviteHandler(message *Message, sql *sql.DB) {
	go func() {
		acceptInviteMsg := Messages.AcceptInviteReq{}
		proto.Unmarshal(message.body, &acceptInviteMsg)
		groupName, err := getGroupNameFromInviteID(acceptInviteMsg.InviteID, sql)
		if err == nil {
			err = addGroupMemberToDatabase(groupName, *message.client.username, sql)
			if err == nil {
				err = deleteInviteFromDatabase(acceptInviteMsg.InviteID, sql)
				if err == nil {
					sendInvites(message.client, sql)
				}
			}
		}
		if err != nil {
			sendError("acceptInviteErr", err, message.client)
		}
	}()
}

func DeleteInviteHandler(message *Message, sql *sql.DB) {
	go func() {
		deleteInviteMsg := Messages.DeleteInviteReq{}
		proto.Unmarshal(message.body, &deleteInviteMsg)
		err := deleteInviteFromDatabase(deleteInviteMsg.InviteID, sql)
		if err != nil {
			sendError("deleteInviteErr", err, message.client)
			return
		}
		sendInvites(message.client, sql)
	}()
}

func getUserInvitesFromDatabase(username string, sql *sql.DB) ([]*Messages.InvitesResp_Invite, error) {
	stmt, _ := sql.Prepare("SELECT id, fromUsername, groupName FROM Invites WHERE toUsername=?")
	rows, execErr := stmt.Query(username)
	if execErr != nil {
		return nil, execErr
	}
	var invites []*Messages.InvitesResp_Invite
	for rows.Next() {
		var inviteID string
		var fromUsername string
		var groupName string
		rows.Scan(&inviteID, &fromUsername, &groupName)
		invite := &Messages.InvitesResp_Invite{
			InviteID: inviteID,
			FromUsername: fromUsername,
			GroupName: groupName,
		}
		invites = append(invites, invite)
	}
	return invites, nil
}

//Send all invites where the user is receipient
func sendInvites(client *Client, sql *sql.DB) {
	invites, err := getUserInvitesFromDatabase(*client.username, sql)
	if err != nil {
		sendError("getInvitesErr", err, client)
		return
	}

	invitesMsg := Messages.InvitesResp{Invites: invites}
	invitesData, _ := proto.Marshal(&invitesMsg)
	client.send("getInvites", invitesData)
}

func GetInvitesHandler(message *Message, sql *sql.DB) {
	go sendInvites(message.client, sql)
}