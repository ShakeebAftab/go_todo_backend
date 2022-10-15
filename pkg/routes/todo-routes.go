package routes

import (
	"github.com/gorilla/mux"
)

var RegisterTodoRoutes = func (router *mux.Router) {
	router.HandleFunc("/todos/", controllers.createTodo).Methods("POST") // Create a new todo
	router.HandleFunc("/todos/", controllers.getAllTodos).Methods("GET") // Get all todos
	router.HandleFunc("/todos/{todoId}", controllers.getSingleTodo).Methods("GET") // Get a single todo
	router.HandleFunc("/todos/{todoId}", controllers.updateTodo).Methods("PATCH") // Update a single todo
	router.HandleFunc("/todos/{todoId}", controllers.deleteTodo).Methods("DELETE") // Delete a single todo
}