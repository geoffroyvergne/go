package main

import (
	"log"
	"strings"
	"github.com/firstrow/tcp_server"
)

func main() {

	server := tcp_server.New("localhost:8000")

	server.OnNewClient(func(c *tcp_server.Client) {		
		// c.Send("Hello\n")
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {		
		message = strings.TrimRight(message, "\r\n")

		if message != "" {
			log.Println(message)
		}

		if message == "quit" {
			log.Println("Bye")	
			c.Close()
		}

		c.SendBytes(([]byte(message + " received \n")))
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		log.Println("client gone")
	})

	server.Listen()
}

