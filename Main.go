package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

import _ "github.com/go-sql-driver/mysql"

type MsgHandler func(message *Message, sql *sql.DB)

var clients = map[*Client]bool{}
var msgHandlers = map[string]MsgHandler{}

//Run client loops and add to map
func initClient(client *Client) {
	clients[client] = true
	go client.runSend()
	go client.runRead()
}

//Single handler for all messages
func handleMessage(message *Message, sql *sql.DB) {
	handler, containsHandler := msgHandlers[message.typeID]
	if containsHandler {
		handler(message, sql)
	} else {
		log.Println("No type handler for id: ", message.typeID)
	}
}

//Cleanup client on disconnect
func deleteClient(client *Client, sql *sql.DB) {
	delete(clients, client)

	leaveActiveGroup(client, sql)

	err := (*client.connection).Close()
	if err != nil {
		log.Println("Connection Close Error: ", err)
	}
}

//Handle all client events on a single loop
func runNetEvents(connectChannel chan *Client, recvMsgChannel chan *Message, disconnectChannel chan *Client, sql *sql.DB) {
	for {
		select {
			case msg, more := <- recvMsgChannel:
				if !more {
					return
				}
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
				deleteClient(client, sql)
		}
	}
}

//Read in configuration from JSON file
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

//Link handlers to message types
func linkHandlers() {
	//UserManager
	msgHandlers["signUp"] = CreateAccountHandler
	msgHandlers["login"] = LoginHandler
	msgHandlers["searchUsers"] = SearchHandler
	//InviteManager
	msgHandlers["invite"] = InviteToGroupHandler
	msgHandlers["getInvites"] = GetInvitesHandler
	msgHandlers["acceptInvite"] = AcceptInviteHandler
	msgHandlers["deleteInvite"] = DeleteInviteHandler
	//GroupManager
	msgHandlers["createGroup"] = CreateGroupHandler
	msgHandlers["getGroups"] = GetGroupsHandler
	msgHandlers["joinGroup"] = JoinGroupHandler
	msgHandlers["leaveGroup"] = LeaveGroupHandler
	msgHandlers["refreshGroup"] = RefreshGroupHandler
	//MessageManager
	msgHandlers["textMsg"] = TextMessageHandler
	//FileManager
	msgHandlers["upload"] = UploadHandler
	msgHandlers["download"] = DownloadHandler
}

//Create client and send to the net loop for initialization
func onConnection(connection *tls.Conn, connectChannel chan *Client, recvMsgChannel chan *Message, disconnectChannel chan *Client) {
	client := &Client{
		connection: connection,
		sendChannel: make(chan *Message),
		recvChannel: recvMsgChannel,
		disconnectChannel: disconnectChannel,
		username: nil,
	}
	connectChannel <- client
}

func runServer(config *Configuration, db *sql.DB) {
	connectChannel := make(chan *Client)
	recvMsgChannel := make(chan *Message)
	disconnectChannel := make(chan *Client)
	go runNetEvents(connectChannel, recvMsgChannel, disconnectChannel, db)
	cer, err := tls.LoadX509KeyPair("./tls/localhost.crt", "./tls/localhost.key")
	if err != nil {
		log.Fatal(err)
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

	l, err := tls.Listen("tcp", config.ListenAddress + ":" + config.ListenPort, tlsConfig)
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	for {
		log.Println("LISTENING FOR CONNECTION")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		tlsConn, ok := conn.(*tls.Conn)
		if !ok {
			fmt.Println("UPGRADE FAILED")
		}
		handshakeErr := tlsConn.Handshake()
		if handshakeErr != nil {
			fmt.Println("Handshake err: ", handshakeErr)
		}
		onConnection(tlsConn, connectChannel, recvMsgChannel, disconnectChannel)
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