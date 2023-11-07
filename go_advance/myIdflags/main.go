package main

import (
	auth "demo/myIdflags/auth"
	todo "demo/myIdflags/todo"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {

	err := godotenv.Load("local.env")
	if err != nil {
		log.Panicln("Please consider env var: %s", err)
	}

	db, err := gorm.Open(sqlite.Open("som.db"), &gorm.Config{})
	if err != nil {
		panic("failed")
	}

	err = db.AutoMigrate(&todo.Todo{})
	if err != nil {
		return
	}
	r := gin.Default()
	r.GET("/x", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})
	r.GET("/tokenz", auth.AccessToken(os.Getenv("SIGN")))
	protected := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))

	handler := todo.NewTodoHandler(db)
	protected.POST("/todos", handler.NewTask)

	r.Run()
}
