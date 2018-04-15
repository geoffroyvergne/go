package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
	testapplib "testapp/lib"
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

	router.GET("/" + prefix + "/testapp/localvalue", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": testapplib.GetLocalValue(),
		})
	})
	
	router.Run(host)
}
