package todo

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	auth "github.com/s7010390/myDemo/auth"
	"gorm.io/gorm"
)

type Todo struct {
	title string `json:"text" binding:"required"`
	gorm.Model
}

func (Todo) TableName() string {
	return "todo"
}

type TodoHandler struct {
	db *gorm.DB
}

func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {

	s := c.Request.Header.Get("Authorization")
	tokenSting := strings.TrimPrefix(s, "Bearer ")

	if err := auth.Protect(tokenSting); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID,
	})

}
