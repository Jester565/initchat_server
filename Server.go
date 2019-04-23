package main

import (
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

type Lobby struct {
	clients   []*Client
	chatRooms map[string]*ChatRoom
}

func NewLobby() *Lobby {
	lobby := &Lobby{
		clients:   make([]*Client, 0),
		chatRooms: make(map[string]*ChatRoom),
	}

	return lobby
}

func (lobby *Lobby) Join(client *Client) {
	lobby.clients = append(lobby.clients, client)
	client.Write("Welcome to GoChat, you'll need an account to continue\n")
	client.authenticate()
	go lobby.start(client)
}

func (lobby *Lobby) start(client *Client) {
	client.menu()
	for {
		str, _ := client.reader.ReadString('\n')
		lobby.ParseLobby(str, client)
	}
}

//used to get username to test other methods
func (client *Client) authenticate() {
	client.Write("Username: \n")
	name, _ := client.reader.ReadString('\n')
	client.UserName = strings.TrimSpace(name)
}

func (lobby *Lobby) Read(client *Client) {
	for {
		str, _ := client.reader.ReadString('\n')
		lobby.ParseLobby(str, client)
	}
}

func (lobby *Lobby) CreateChatRoom(client *Client, name string) {
	chatRoom := NewChatRoom(name)
	lobby.chatRooms[name] = chatRoom
	client.ChatRoomCommands(name)
	lobby.JoinChatRoom(client, name)
}

func (client *Client) ChooseChatRoomName() string {
	client.Write("Please pick a name for your chat room.\n")
	name, _ := client.reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func (lobby *Lobby) JoinChatRoom(client *Client, name string) {
	if lobby.chatRooms[name] == nil {
		client.Write("Incorrect. Try another name.\n")
		lobby.ParseLobby("2", client) //sends second option back to the menu so user can try again
	}

	lobby.chatRooms[name].Join(client, lobby)
	client.chatRoom.Send(client.UserName + " has joined.")
}

func (lobby *Lobby) AvailableChatRooms(client *Client) {
	client.Write("Available Chat Rooms:\n")
	for name := range lobby.chatRooms {
		client.Write(name)
		client.Write("\n")
	}
}

func (lobby *Lobby) getName(client *Client) string {
	client.Write("Please choose a chat room.\n")
	name, _ := client.reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func (lobby *Lobby) ParseLobby(message string, client *Client) {
	switch {
	case strings.TrimSpace(message) == "1":
		name := client.ChooseChatRoomName()
		lobby.CreateChatRoom(client, name)
	case strings.TrimSpace(message) == "2":
		if len(lobby.chatRooms) == 0 {
			client.Write("No chat rooms.\n")
			client.Write("Choose a different option.\n")
		} else {
			lobby.AvailableChatRooms(client)
			name := lobby.getName(client)
			lobby.JoinChatRoom(client, name)
			client.ChatRoomCommands(client.chatRoom.name)
		}
	case strings.TrimSpace(message) == "3":
		//view invites
	default:
		client.Write("Please choose a correct menu option.\n")
	}

}

func (client *Client) menu() {
	client.Write("1. Create a Chat Group\n")
	client.Write("2. Open a Previous Chat Group\n")
	client.Write("3. View invites\n")
}

type ChatRoom struct {
	name     string
	clients  []*Client
	messages []string
	invites  []string
}

func NewChatRoom(name string) *ChatRoom {
	return &ChatRoom{
		name:     name,
		clients:  make([]*Client, 0),
		messages: make([]string, 0),
		invites:  make([]string, 0),
	}
}

func (client *Client) ChatRoomCommands(name string) {
	client.Write("---------- " + name + " Group " + "----------\n")
	client.Write("Commands: ~invite  #Invite a user\n" +
		"  \t\t~leave #Leave the group\n" +
		"  \t\t~fs #Send file\n" +
		"  \t\t~dwnld #Download file\n")
}

func (chatRoom *ChatRoom) Join(client *Client, lobby *Lobby) {
	client.chatRoom = chatRoom
	for _, message := range chatRoom.messages {
		client.Write(message)
	}
	chatRoom.clients = append(chatRoom.clients, client)
	for {
		chatRoom.ReadChatRoom(client, lobby)
	}

}

func (chatRoom *ChatRoom) ReadChatRoom(client *Client, lobby *Lobby) {
	for {
		str, _ := client.reader.ReadString('\n')
		chatRoom.ParseChatRoom(str, client, lobby)
	}
}

func (chatRoom *ChatRoom) LeaveChatRoom(client *Client, lobby *Lobby) {
	client.chatRoom.Send(client.UserName + " is leaving the chat room.\n")
	client.chatRoom = nil
	lobby.start(client)
}

func (chatRoom *ChatRoom) ParseChatRoom(message string, client *Client, lobby *Lobby) {
	switch {
	case strings.HasPrefix(message, "~invite"):
		client.Write("test\n")
	case strings.TrimSpace(message) == "~leave":
		client.Write("\n")
		chatRoom.LeaveChatRoom(client, lobby)
	case strings.HasPrefix(message, "~fs"):
		//send file
	case strings.HasPrefix(message, "~dwnld"):
		//download file
	default:
		appendMessage := time.Now().Format("2006-01-02 15:04:05") + " " + client.UserName + " >> " + message
		client.chatRoom.Send(appendMessage)
	}

}

func (chatRoom *ChatRoom) Send(message string) {
	chatRoom.messages = append(chatRoom.messages, message)
	for _, client := range chatRoom.clients {
		client.Write(message)
	}
}

/*func (lobby *Lobby)Invite(person string, chatRoom *ChatRoom, client *Client){
	client.Write("All connected clients:\n")
	for p := range lobby.clients {
		client.Write(p)
		client.Write("\n")
	}
}
*/

type Client struct {
	UserName string
	chatRoom *ChatRoom
	conn     net.Conn
	reader   *bufio.Reader
	writer   *bufio.Writer
}

func NewClient(conn net.Conn) *Client {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)

	client := &Client{
		UserName: "Client",
		chatRoom: nil,
		conn:     conn,
		reader:   reader,
		writer:   writer,
	}
	return client
}

func (client *Client) Write(str string) {
	client.writer.WriteString(str)
	client.writer.Flush()
}

func main() {
	lobby := NewLobby()
	listener, _ := net.Listen("tcp", ":8081")
	log.Println("Listening on " + ":8081")

	for {
		conn, _ := listener.Accept()
		lobby.Join(NewClient(conn))
	}
}
