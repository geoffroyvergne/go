package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	host := flag.String("h", "localhost:8000", "host:port ex : localhost:8000")
	directory := flag.String("d", ".", "the directory of static file to host")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on : %s\n", *directory, *host)
	log.Fatal(http.ListenAndServe(*host, nil))

	//fs := http.FileServer(http.Dir(*directory))
	//http.Handle("/", http.StripPrefix("/", fs))
	//http.ListenAndServe(*host, nil)
}
