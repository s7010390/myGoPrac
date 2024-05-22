package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/natchapon/todoapi/router"
	"github.com/natchapon/todoapi/store"
	"github.com/natchapon/todoapi/todo"
)

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}

	r := router.NewFiberRouter()

	//gormStore := store.NewGormStore()
	mongoStore := store.NewMongoDBStore()

	handler := todo.NewTodoHandler(mongoStore)
	r.POST("/todos", handler.NewTask)

	if err := r.Listen(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
