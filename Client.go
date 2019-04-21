package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func Read(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Disconnected from server.\n")
			wg.Done()
			return
		}
		fmt.Print(str)
	}
}

func Write(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(conn)

	for {
		str, _ := reader.ReadString('\n')
		writer.WriteString(str)
		writer.Flush()

	}
}

func main() {
	wg.Add(1)
	conn, _ := net.Dial("tcp", ":8081")
	go Write(conn)
	go Read(conn)
	wg.Wait()
}
