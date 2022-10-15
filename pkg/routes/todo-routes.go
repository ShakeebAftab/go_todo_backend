package routes

import (
	"github.com/gorilla/mux"
	"github.com/shakeebaftab/todo_backend/pkg/controllers"
)

var RegisterTodoRoutes = func (router *mux.Router) {
	router.HandleFunc("/todos/", controllers.CreateTodo).Methods("POST") // Create a new todo
	router.HandleFunc("/todos/", controllers.GetAllTodos).Methods("GET") // Get all todos
	router.HandleFunc("/todos/{todoId}", controllers.GetSingleTodo).Methods("GET") // Get a single todo
	router.HandleFunc("/todos/{todoId}", controllers.UpdateTodo).Methods("PUT") // Update a single todo
	router.HandleFunc("/todos/{todoId}", controllers.DeleteTodo).Methods("DELETE") // Delete a single todo
}