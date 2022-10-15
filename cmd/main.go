package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shakeebaftab/todo_backend/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterTodoRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}