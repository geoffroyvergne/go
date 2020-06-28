package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
)

func printHelp() string {
	return "Avaliable commands : ping help quit\n"
}

func handle(conn net.Conn) {
	defer conn.Close()

	//content, err := Read(con)
	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	
	if err != nil {
		log.Printf("Listener: Exit!")
		conn.Close()
		return
	}

	//writer, err := bufio.NewWriter(conn);

	message := string(bufferBytes)
	message = strings.Replace(message, "\r\n", "", -1)
	//clientAddr := conn.RemoteAddr().String()
	//response := "Message : " + message + " from " + clientAddr + "\n"

	//log.Printf("New command: %s from %s \n", message, conn.RemoteAddr().String())

	switch message {
		case "ping": 
			conn.Write([]byte(string("pong\n")))
		case "help":	
			conn.Write([]byte(printHelp()))
		case "quit": 
			conn.Write([]byte(string("bye\n")))
			conn.Close()
			log.Println("client exit from " + conn.RemoteAddr().String())
			return
		
		default: 
			conn.Write([]byte(string("unknown command : " + message + "\n")))
	}

	//fmt.Printf(response)

	handle(conn)
}

func main() {
	tcpPort := ":8888"
	fmt.Println("Tcp server listen on " + tcpPort + " ... ")

	listener, err := net.Listen("tcp", tcpPort)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		//break
		log.Println("new client from " + conn.RemoteAddr().String())
		conn.Write([]byte(string("Welcome ! " + printHelp())))

		go handle(conn)
	}
}
