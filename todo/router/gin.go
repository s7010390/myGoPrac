package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	todo "github.com/s7010390/myTodo/todo"
)

type MyContext struct {
	*gin.Context
}

func NewMyContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c}
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *MyContext) JSON(statuscode int, v interface{}) {
	c.Context.JSON(statuscode, v)
}

func (c *MyContext) TransactionID() string {
	return c.Context.Request.Header.Get("transactionID")
}

func (c *MyContext) Audiance() string {
	if aud, ok := c.Context.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func NewGinHandler(handler func(todo.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewMyContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}
	config.AllowHeaders = []string{
		"origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))
	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(todo.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}
