package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	writer   *bufio.Writer
	reader   *bufio.Reader
	conn     net.Conn
	UserName string
}

func newClient(conn net.Conn) {
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	client := &Client{UserName: "UserExample",
		writer: writer,
		reader: reader,
	}
	client.handler(conn)
}

func (client *Client) handler(conn net.Conn) {
	for {
		message, _ := client.reader.ReadString('\n')
		fmt.Print("Message Received:", string(message))
		appendMessage := client.UserName + " >> " + message
		//client.incoming <- message
		client.writer.WriteString(appendMessage)
		client.writer.Flush()
	}
}

func getUserName() string {
	fmt.Println("Please choose a username.")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Print(name)
	return name

}

func main() {

	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8081")

	for {
		conn, _ := ln.Accept()
		newClient(conn)
	}
}
