package client

import(
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
	"time"
	//"io"
)

// TCPClient TCP Client
func TCPClient(host string, text string, timeout time.Duration) {

	if text == "" {
		fmt.Println("text requiered")
		os.Exit(0)
	}

	conn, _ := net.DialTimeout("tcp", host, timeout *time.Second)

	handleRequest(conn, text)
	defer conn.Close()
	
	//os.Exit(0)
}

func handleRequest(conn net.Conn, text string) {
	//defer conn.Close()
	
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Text to send: ")
	//text, _ := reader.ReadString('\n')

	//conn.Write([]byte(text))
	//bufio.NewWriter(conn).Write([]byte(text))

	//io.WriteString(conn, text)

	fmt.Fprintf(conn, text + "\n")
	message, _, _ := bufio.NewReader(conn).ReadLine()
	log.Println("Message from server: " + string(message))

	//conn.Close()
}