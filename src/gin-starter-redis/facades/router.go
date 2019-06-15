package facades

import (
	"github.com/gin-gonic/gin"
)

func Router(prefix string, router *gin.Engine ) {
	base := router.Group("/" + prefix)
	{
		base.GET("/", Index)
		base.GET("/index/test", IndexTest)

		base.GET("/debug/echo", DebugEcho)
		base.GET("/debug/ping", DebugPing)
		base.GET("/debug/ping/redis", DebugPingRedis)

		base.GET("/todo/:todo", TodoGet)
		base.POST("/todo/:todo", TodoPost)
		base.DELETE("/todo/:todo", TodoDelete)
	}
}