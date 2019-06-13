package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test() string {
	return "Hello from this external package"
}

func Localvalue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "localvalue",
	})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}