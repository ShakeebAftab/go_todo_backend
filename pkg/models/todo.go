package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shakeebaftab/todo_backend/pkg/config"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	id int64 `gorm:"primaryKey;autoIncrement"` 
	content string
	isCompleted bool `gorm:"default:false"` 
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todo{})
}

func (todo *Todo) CreateTodo() *Todo {
	db.NewRecord(todo)
	db.Create(&todo)
	return todo
}

func GetAllTodos() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func GetSingleTodo(id int64) *Todo {
	var todo *Todo
	db.First(&todo, id)
	return todo
}

func UpdateTodo(updatedTodo *Todo) *Todo {
	var todo *Todo = GetSingleTodo(updatedTodo.id)
	todo = updatedTodo
	db.Save(todo)
	return todo
}

func DeleteTodo(id int64) string {
	db.Delete(&Todo{}, id)
	return "Todo removed successfully"
}