package router

import (
	"todo-odev/pkg/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todos/{id}/complete", handlers.MarkAsCompleted).Methods("PUT")

	return r
}
