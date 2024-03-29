package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Todo struct {
	Author string `json:"author"`
	Title  string `json:"text" binding:"required"`

	gorm.Model
}

func (Todo) TableName() string {
	return "todos"
}

type storer interface {
	MyNew(*Todo) error
	MyFind(*[]Todo) error
	MyDelete(*Todo, int) error
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TransactionID() string
	Audiance() string
}

type TodoHandler struct {
	store storer
}

func NewTodoHandler(store storer) *TodoHandler {
	return &TodoHandler{store: store}
}

func (t *TodoHandler) NewTask(c Context) {

	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//loging
	if todo.Title == "sleep" {
		transactionID := c.TransactionID()
		aud := c.Audiance()
		log.Println(transactionID, aud, "not allowed")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not allowed",
		})
		return
	}
	err := t.store.MyNew(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID":     todo.Model.ID,
		"author": todo.Author,
	})

}
func (t *TodoHandler) List(c *gin.Context) {
	var todos []Todo
	err := t.store.MyFind(&todos)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) Remove(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := t.store.MyDelete(&Todo{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
