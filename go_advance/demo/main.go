package main

import (
	auth "github.com/s7010390/myDemo/auth"
	todo "github.com/s7010390/myDemo/todo"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("som.db"), &gorm.Config{})
	if err != nil {
		panic("failed")
	}
	db.AutoMigrate(&todo.Todo{})
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tokenz", auth.AccessToken)

	handler := todo.NewTodoHandler(db)
	r.POST("/todos", handler.NewTask)
	r.Run()
}
