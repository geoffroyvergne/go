package server

import(
	"fmt"
	"net"
	"bufio"
	"strings"
)

// TCPServer TCP Server
func TCPServer(host string) {
	fmt.Println("Launching server on " + host + "...")

	ln, _ := net.Listen("tcp", host)
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		newmessage := strings.ToUpper(message)
		conn.Write([]byte(newmessage + "\n"))
	  }
}
