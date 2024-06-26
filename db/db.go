package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createTodosTable := `
	CREATE TABLE IF NOT EXISTS todos (
		todoId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		text TEXT NOT NULL,
		completed BOOLEAN DEFAULT 0 
	)
	`
	_, err := DB.Exec(createTodosTable)
	if err != nil {
		panic("could not create todo table")
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		userId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
	)
	`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		panic("could not create users table")
	}
}
