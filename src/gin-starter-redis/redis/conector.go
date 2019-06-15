package redis

import (
	"github.com/go-redis/redis"
	"log"
)

// RedisInstance main client var
var RedisInstance *redis.Client

// Connect connection redis client
func Connect(host string, password string, db int) {

	//done := make(chan bool)

	//go func() {
		RedisInstance = redis.NewClient(&redis.Options{
			Addr:     host, //"localhost:6379",
			Password: password, //"", // no password set
			DB:       db,  // use default DB
		})
	//}()

	//<-done

	pong, err := Pong()

	if err != nil {
		log.Fatalf("redis: %s\n", err)
	} else {
		log.Printf("Redis ping result : " + pong)
	}
	
	//defer RedisInstance.Close()
}

// Pong func
func Pong() (string, error) {
	return RedisInstance.Ping().Result()
}