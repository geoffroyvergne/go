package main

import(
	"flag"
	"fmt"
	"os"
	client "net/simpletcp/client"
	server "net/simpletcp/server"
)

func main() {
	tcpType := flag.String("t", "", "type : client or server")
	host := flag.String("h", "localhost:8000", "host:port ex : localhost:8000")

	flag.Parse()

	if(*tcpType == "") {
		fmt.Println("type requiered")
		os.Exit(0)
	}

	fmt.Println("TCP main " + *tcpType)

	if (*tcpType == "client") {
		client.TCPClient(*host)
	} else if (*tcpType == "server") {
		server.TCPServer(*host)
	} else {
		fmt.Println("unknown tcp type " + *tcpType)
	}
}
