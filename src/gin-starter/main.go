package main

import (
	"github.com/gin-gonic/gin"
	"gin-starter/facades"
	"os"
	"log"
)

func main() {
	//Init("api", "localhost:3000")

	prefix := "api"
	if tmpPrefix := os.Getenv("API_PREFIX"); tmpPrefix != "" {
		prefix = tmpPrefix
	}

	host := "localhost:3000"
	if tmpHost := os.Getenv("API_HOST"); tmpHost != "" {
		host = tmpHost
	}

	Init(prefix, host)
}

func Init(prefix, host string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	facades.Router(prefix, router)

	// golang closing channels http server
	
	done := make(chan bool)

	go func() {
		if err := router.Run(host); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Server started %s", host)

	<-done
}
