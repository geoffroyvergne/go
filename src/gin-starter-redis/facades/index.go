package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func Index(c *gin.Context) {
	log.Println("index")

	c.JSON(http.StatusOK, gin.H{
		"message": "index",
		"status": "200",
	})
}

func IndexTest(c *gin.Context) {
	log.Println("index/test")

	c.JSON(http.StatusOK, gin.H{
		"message": "index test",
		"status": "200",
	})
}