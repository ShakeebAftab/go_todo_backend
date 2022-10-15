package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shakeebaftab/todo_backend/pkg/models"
	"github.com/shakeebaftab/todo_backend/pkg/utils"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := &models.Todo{}
	utils.ParseBody(r, newTodo)
	todo := newTodo.CreateTodo()
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()
	res, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSingleTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	id, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing todo id")
	}
	todo := models.GetSingleTodo(id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	newTodo := &models.Todo{}
	utils.ParseBody(r, newTodo)
	todo := models.UpdateTodo(newTodo)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	id, err := strconv.ParseInt(todoId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing todo id")
	}
	todo := models.DeleteTodo(id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}