package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

import _ "github.com/go-sql-driver/mysql"

type MsgHandler func(message *Message, sql *sql.DB)

var clients = map[*Client]bool{}
var msgHandlers = map[string]MsgHandler{}

func onConnection(connection *net.Conn, connectChannel chan *Client, recvMsgChannel chan *Message, disconnectChannel chan *Client) {
	client := &Client{
		connection: connection,
		sendChannel: make(chan *Message),
		recvChannel: recvMsgChannel,
		disconnectChannel: disconnectChannel,
		username: nil,
	}
	connectChannel <- client
}

func initClient(client *Client) {
	clients[client] = true
	go client.runSend()
	go client.runRead()
	log.Println("CLIENT INTIALIZED")
}

func handleMessage(message *Message, sql *sql.DB) {
	handler, containsHandler := msgHandlers[message.typeID]
	if containsHandler {
		handler(message, sql)
	} else {
		log.Println("No type handler for id: ", message.typeID)
	}
}

func disconnectClient(client *Client) {
	delete(clients, client)

	leaveActiveGroup(client)

	err := (*client.connection).Close()
	if err != nil {
		log.Println("Connection Close Error: ", err)
	}
}

func runNetEvents(connectChannel chan *Client, recvMsgChannel chan *Message, disconnectChannel chan *Client, sql *sql.DB) {
	for {
		select {
			case msg, more := <- recvMsgChannel:
				if !more {
					return
				}
				log.Println("RUN MESSAGE: ", msg.typeID)
				handleMessage(msg, sql)
			case client, more := <- connectChannel:
				if !more {
					return
				}
				initClient(client)
			case client, more := <- disconnectChannel:
				if !more {
					return
				}
				disconnectClient(client)
		}
	}
}

func linkHandlers() {
	msgHandlers["signUp"] = createAccountHandler
	msgHandlers["login"] = loginHandler
	msgHandlers["createGroup"] = createGroupHandler
	msgHandlers["getGroups"] = getGroupsHandler
	msgHandlers["searchUsers"] = searchHandler
	msgHandlers["invite"] = inviteToGroupHandler
	msgHandlers["getInvites"] = getInvitesHandler
	msgHandlers["acceptInvite"] = acceptInviteHandler
	msgHandlers["deleteInvite"] = deleteInviteHandler
	msgHandlers["joinGroup"] = joinGroupHandler
	msgHandlers["textMsg"] = textMessageHandler
	msgHandlers["leaveGroup"] = leaveGroupHandler
	msgHandlers["refreshGroup"] = refreshGroupHandler
}

type Configuration struct {
	ListenAddress string
	ListenPort string
	DbAddress string
	DbPort string
	DbUser string
	DbPwd string
	DbDb string
}

func readConfig(filename string) (*Configuration, error) {
	file, fErr := os.Open(filename)
	if fErr != nil {
		log.Println("File open error: ", fErr)
		return nil, fErr
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}

func runServer(config *Configuration, db *sql.DB) {
	connectChannel := make(chan *Client)
	recvMsgChannel := make(chan *Message)
	disconnectChannel := make(chan *Client)
	go runNetEvents(connectChannel, recvMsgChannel, disconnectChannel, db)

	l, err := net.Listen("tcp", "0.0.0.0:2750")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	for {
		log.Println("LISTENING FOR CONNECTION")
		connection, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		onConnection(&connection, connectChannel, recvMsgChannel, disconnectChannel)
		log.Println("CONNECTION ACCEPTED")
	}
}

func main() {
	linkHandlers()

	config, err := readConfig("config.json")
	if err != nil {
		log.Fatalln("Could not read config: ", err)
		return
	}
	dbConStr := config.DbUser + ":" + config.DbPwd + "@tcp(" + config.DbAddress + ":" + config.DbPort + ")/" + config.DbDb
	db, err := sql.Open("mysql", dbConStr)
	if err != nil {
		log.Fatalln("Could not connect to database: ", err)
		return
	}
	runServer(config, db)
}