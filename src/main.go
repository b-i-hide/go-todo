package main

import (
	//"localhost/go-todo"
	"github.com/gin-gonic/gin"
	"localhost/go-todo/db"
	"localhost/go-todo/src/controllers"
	"fmt"
	"time"
)

func main() {
	r := gin.Default()
	db.DbInit()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello world")
	})
	r.GET("/huga", func(c *gin.Context) {
		c.String(200, "fuga")
	})

	r.GET("/todos", func(c *gin.Context) {
		ctrl := controllers.NewTodo()
		todos := ctrl.GetAll()
		fmt.Println(todos)
		c.JSON(200, gin.H{
			"todos": todos,
		})
	})
	r.POST("/todos", func(c *gin.Context) {
		ctrl := controllers.NewTodo()
		ctrl.Create("hideaki", false, time.Date(2018, 5, 31, 0, 0, 0, 0, time.Local), time.Now())

	})

	r.Run(":8080")

	//b := go_todo.New()
	//b.Init()
}

