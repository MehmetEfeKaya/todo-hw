package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-odev/pkg/database"
	"todo-odev/pkg/models"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query("SELECT * FROM todos")
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		rows.Scan(&todo.ID, &todo.Task, &todo.Completed)
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	statement, _ := database.DB.Prepare("INSERT INTO todos (task, completed) VALUES (?, ?)")
	statement.Exec(todo.Task, false)

	w.WriteHeader(http.StatusCreated)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	statement, _ := database.DB.Prepare("DELETE FROM todos WHERE id = ?")
	statement.Exec(id)

	w.WriteHeader(http.StatusOK)
}

func MarkAsCompleted(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var completed bool
	database.DB.QueryRow("SELECT completed FROM todos WHERE id = ?", id).Scan(&completed)

	statement, _ := database.DB.Prepare("UPDATE todos SET completed = ? WHERE id = ?")
	statement.Exec(!completed, id)

	w.WriteHeader(http.StatusOK)
}
