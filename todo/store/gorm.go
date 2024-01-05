package store

import (
	todo "github.com/s7010390/myTodo/todo"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (store *GormStore) MyNew(todo *todo.Todo) error {
	return store.db.Create(todo).Error
}

func (store *GormStore) MyFind(todos *[]todo.Todo) error {
	return store.db.Find(todos).Error

}
func (store *GormStore) MyDelete(todo *todo.Todo, i int) error {
	return store.db.Delete(todo, i).Error
}
