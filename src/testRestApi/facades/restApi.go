package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Init rest API
func Init(prefix, host string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	
	router.GET("/" + prefix + "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// ping(router)
	
	router.Run(host)
}
