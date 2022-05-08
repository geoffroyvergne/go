package main

import (
    "net/http"
	"log"
	"io"
    "github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//w.Write([]byte("Gorilla!\n"))
	io.WriteString(w, `{"message": "Gorilla"}`)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", TestHandler)

    log.Fatal(http.ListenAndServe(":8000", r))
}