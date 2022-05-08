package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"gin-starter-redis/models"
	"gin-starter-redis/redis"
)

func DebugEcho(c *gin.Context) {
	log.Println("debug/echo")

	status := models.Status
	status.Message = "echo"
	status.Code = 200

	c.JSON(http.StatusOK, status)
}

func DebugPing(c *gin.Context) {
	log.Println("debug/ping")

	status := models.Status
	status.Message = "ping"
	status.Code = 200

	c.JSON(http.StatusOK, status)
}

func DebugPingRedis(c *gin.Context) {
	log.Println("debug/ping/redis")

	pong, _ := redis.RedisInstance.Ping().Result()

	log.Println("ping : " + pong)

	status := models.Status
	status.Message = pong
	status.Code = 200

	c.JSON(http.StatusOK, status)
}