package main

import(
	"flag"
	"os"
	"log"
	client "net/simpletcp/client"
	server "net/simpletcp/server"
)

func main() {
	tcpType := flag.String("t", "", "type : client or server")
	host := flag.String("h", "localhost:8000", "host:port ex : localhost:8000")
	text := flag.String("m", "", "message to send")

	flag.Parse()

	if *tcpType == "" {
		log.Println("type requiered")
		os.Exit(0)
	}

	log.Println("TCP main " + *tcpType)

	if *tcpType == "client" {
		client.TCPClient(*host, *text, 5)
	} else if *tcpType == "server" {
		server.TCPServer(*host)
	} else {
		log.Println("unknown tcp type " + *tcpType)
		os.Exit(0)
	}
}
