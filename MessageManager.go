package main

import (
	"./Messages"
	"database/sql"
	"github.com/golang/protobuf/proto"
	"time"
)

func getGroupMessages(groupName string, sql *sql.DB) ([]*Messages.TextMessage, error) {
	stmt, _ := sql.Prepare("SELECT creator, contents, creationTime FROM Messages WHERE groupName=? ORDER BY creationTime")
	rows, execErr := stmt.Query(groupName)
	if execErr != nil {
		return nil, execErr
	}
	var messages []*Messages.TextMessage
	for rows.Next() {
		var creator string
		var contents string
		var creationTime uint64
		rows.Scan(&creator, &contents, &creationTime)
		messages = append(messages, &Messages.TextMessage{
			Username: creator,
			Message: contents,
			Time: creationTime,
		})
	}
	return messages, nil
}

func addGroupMessageToDatabase(contents string, creationTime uint64, groupName string, username string, sql *sql.DB) error {
	stmt, _ := sql.Prepare("INSERT INTO Messages VALUES (?, ?, ?, ?)")
	_, execErr := stmt.Exec(groupName, username, contents, creationTime)
	if execErr != nil {
		return execErr
	}
	return nil
}

func sendGroupMessage(contents string, creationTime uint64, groupName string, username string) error {
	activeGroupsMutex.RLock()
	defer activeGroupsMutex.RUnlock()
	group, hasGroup := activeGroups[groupName]
	if hasGroup {
		textMsg := Messages.TextMessage{
			Username: username,
			Message: contents,
			Time: creationTime,
		}
		textMsgData, _ := proto.Marshal(&textMsg)
		for client, _ := range group.clients {
			if *client.username != username {
				client.send("message", textMsgData)
			}
		}
	}
	return nil
}

func addGroupMessage(contents string, groupName string, username string, db *sql.DB) {
	creationTime := uint64(time.Now().Unix())
	go addGroupMessageToDatabase(contents, creationTime, groupName, username, db)
	go sendGroupMessage(contents, creationTime, groupName, username)
}

func TextMessageHandler(message *Message, sql *sql.DB) {
	textMsg := Messages.TextMessageReq{}
	proto.Unmarshal(message.body, &textMsg)
	addGroupMessage(textMsg.Message, *message.client.groupName, *message.client.username, sql)
}