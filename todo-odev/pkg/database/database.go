package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) {
	DB, _ = sql.Open("sqlite3", filepath)
	createTable()
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,        
        "task" TEXT,
        "completed" BOOLEAN
    );`

	statement, _ := DB.Prepare(createTableSQL)
	statement.Exec()
}
