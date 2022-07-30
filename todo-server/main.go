package main

import (
	"net/http"
	"strconv"

	"example.com/todo-server/db"
	"example.com/todo-server/types"
	"github.com/gin-gonic/gin"
)



func main() {
	var service db.DbServiceInterface = db.NewInMemoryDbService()


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/todo")
	})

	r.GET("/todo", func(c *gin.Context) {
		c.JSON(200, service.GetAll())
	})

	r.DELETE("/todo/:id", func(c *gin.Context) {
		var id int
		id, err := strconv.Atoi(c.Param("id"))
		
		if handleGinError(c, err) {
			return
		}

		service.Delete(id)

		c.JSON(200, service.GetAll())
	})

	r.POST("/todo", func(c *gin.Context) {
		var todo types.Todo
		err := c.ShouldBindJSON(&todo)

		if handleGinError(c, err) {
			return
		}

		service.Save(todo)
		c.JSON(200, service.GetAll())
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}



func handleGinError(c *gin.Context, err error) bool {
	if (err != nil) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return true
	}
	return false
}
