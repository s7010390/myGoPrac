package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	auth "github.com/s7010390/myTodo/auth"
	store "github.com/s7010390/myTodo/store"
	todo "github.com/s7010390/myTodo/todo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	/*config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))
	*/
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tokenz", auth.AccessToken(os.Getenv("SIGN")))
	protected := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))
	store := store.NewGormStore(db)
	handler := todo.NewTodoHandler(store)
	protected.POST("/todos", todo.NewGinHandler(handler.NewTask))
	protected.GET("/todos", handler.List)
	protected.DELETE("/todos/:id", handler.Remove)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
	//r.Run()
}
