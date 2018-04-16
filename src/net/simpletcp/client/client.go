package client

import(
	"fmt"
	"net"
	"bufio"
	"os"
)

// TCPClient TCP Client
func TCPClient(host string) {

	fmt.Println("Conected to " + host)

	conn, _ := net.Dial("tcp", host)

	for { 
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text + "\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)
	  }
}
