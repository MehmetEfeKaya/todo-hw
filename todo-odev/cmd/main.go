package main

import (
	"log"
	"net/http"
	"todo-odev/pkg/database"
	"todo-odev/pkg/router"
)

func main() {

	database.InitDB("todos.db")

	r := router.Router()

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
