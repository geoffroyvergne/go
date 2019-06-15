package facades

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"gin-starter-redis/models"
	"gin-starter-redis/redis"
)

func TodoGet(c *gin.Context) {
	todo := c.Param("todo")
	todoResult, _ := redis.RedisInstance.Get(todo).Result()

	log.Println("todo get : " + todo)

	status := models.Status
	status.Message = "todo get" + todo + " " + todoResult
	status.Code = 200

	c.JSON(http.StatusOK, status)
}

func TodoPost(c *gin.Context) {
	todoID := c.Param("todo")
	//todoObject = c.Request.Body
	todoObject := models.Todo

	if err := c.ShouldBindJSON(&todoObject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redis.RedisInstance.Set(todoID, todoObject.Task, 0)

	log.Println("todo post : " + todoID)

	status := models.Status
	status.Message = "todo post : " + todoObject.Task
	status.Code = 200

	c.JSON(http.StatusOK, status)
}

func TodoDelete(c *gin.Context) {

	todoID := c.Param("todo")

	redis.RedisInstance.Del(todoID)

	log.Println("todo delete : " + todoID) 

	status := models.Status
	status.Message = "todo delete : " + todoID
	status.Code = 200

	c.JSON(http.StatusOK, status)
}