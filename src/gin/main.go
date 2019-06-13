package main

import (
	"github.com/gin-gonic/gin"
	"gin/facades"
	//"fmt"
)

func main() {
	Init("api", "localhost:3000")
	//fmt.Println(facades.Test())
}

func Init(prefix, host string) {
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	
	router.GET("/" + prefix + "/ping", facades.Ping)

	router.GET("/" + prefix + "/testapp/localvalue", facades.Localvalue)
	
	router.Run(host)
}
