package main

import (
	"github.com/gin-gonic/gin"
	"gin-starter-redis/redis"
	"gin-starter-redis/facades"
	"os"
	"log"
)

func main() {
	initRedis()
	initAPI()
	defer redis.RedisInstance.Close()
}

func initRedis() {
	redisHost := "localhost:6379"
	if tmpRedisHost := os.Getenv("REDIS_HOST"); tmpRedisHost != "" {
		redisHost = tmpRedisHost
	}

	redis.Connect(redisHost, "", 0)
}

func initAPI() {

	apiPrefix := "api"
	if tmpAPIPrefix := os.Getenv("API_PREFIX"); tmpAPIPrefix != "" {
		apiPrefix = tmpAPIPrefix
	}

	apiHost := "localhost:3000"
	if tmpAPIHost := os.Getenv("API_HOST"); tmpAPIHost != "" {
		apiHost = tmpAPIHost
	}

	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	facades.Router(apiPrefix, router)

	// golang closing channels http server
	
	done := make(chan bool)

	go func() {
		if err := router.Run(apiHost); err != nil {
			log.Fatalf("API listen: %s\n", err)
		}
	}()

	log.Printf("Server started %s", apiHost)

	<-done
}
