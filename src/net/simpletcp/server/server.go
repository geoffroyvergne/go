package server

import(
	"net"
	"bufio"
	"log"
	//"os"
	"io"
)

// TCPServer TCP Server
func TCPServer(host string) {
	log.Println("Launching server on " + host + "...")

	l, err := net.Listen("tcp", host)
    if err != nil {
        log.Fatal(err)
    }
	
	defer l.Close()
	
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }        
		
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	message, _, err := bufio.NewReader(conn).ReadLine()
	messageStr := string(message)

	if err != nil {
		if err != io.EOF {
			log.Printf("Error: %+v", err.Error())
			return
		}
	}

	if messageStr == "" {
		return
	}
	
	if messageStr == "quit" {
		log.Println("quit command received. Bye.")
		conn.Close()
		
		return
	}

	if messageStr == "ping" {
		log.Println("ping command received.")
		conn.Write([]byte("pong\n"))

		return
	}

	log.Println("Message: " + messageStr)
	
	newmessage := messageStr + " from server \n"
	conn.Write([]byte(newmessage))
}