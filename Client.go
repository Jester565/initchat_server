package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)
	for {
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		writer.WriteString(text)
		writer.Flush()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print(message)

	}

}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection established!")
	for {
		handleConnection(conn)
	}
}
