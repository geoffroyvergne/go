package main

import (
	"fmt"
	"net"
	"log"
	"bufio"
	"strings"
	"strconv"
	"io/ioutil"
)

func printHelp() string {
	return "Avaliable commands : ping help quit\n"
}

func handle(conn net.Conn) {

	// GET / HTTP/1.0 <CRLF><CRLF>
    /*
    GET / HTTP/1.1
    Host: localhost:8888
    User-Agent: curl/7.70.0
    Accept: *
	*/
	
	defer conn.Close()

	bufferBytes, err := bufio.NewReader(conn).ReadBytes('\n')
	
	if err != nil {
		log.Printf("Listener: Exit!")
		conn.Close()
		return
	}

	request := string(bufferBytes)
	requestArr := strings.Split(request, " ")
	uri := requestArr[1]
	wwwData := "www";

	content, err := ioutil.ReadFile(wwwData + uri)
	if err != nil {
		content, _ := ioutil.ReadFile(wwwData + "/error/404.html")
		contentLen := strconv.Itoa(len(string(content)))
		response := "HTTP/1.1 404 NOT FOUND\r\nContent-Length: " + contentLen + "\r\n\r\n" + string(content)
		conn.Write([]byte(response))
		log.Println("GET " + uri + " HTTP/1.1 404 NOT FOUND " + conn.RemoteAddr().String());
	} else {
		contentLen := strconv.Itoa(len(string(content)))
		response := "HTTP/1.1 200 OK\r\nContent-Length: " + contentLen + "\r\n\r\n" + string(content)
		conn.Write([]byte(response))
		log.Println("GET " + uri + " HTTP/1.1 200 OK " + conn.RemoteAddr().String());
	}

	//fmt.Printf(response)

	//handle(conn)
	conn.Close();
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
		
		go handle(conn)
	}
}
