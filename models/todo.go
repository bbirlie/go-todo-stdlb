package models

import (
	"database/sql"
	"errors"

	"a.com/http/db"
)

type Todo struct {
	TodoId    int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

func GetAllTodos() ([]Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.TodoId, &todo.Title, &todo.Text, &todo.Completed)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func SelectTodo(id int) (Todo, error) {
	query := "SELECT todoId, title, text, completed FROM todos WHERE todoId = ?"

	row := db.DB.QueryRow(query, id)

	var t Todo

	err := row.Scan(&t.TodoId, &t.Title, &t.Text, &t.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Todo{}, errors.New("models: no matching record found")
		} else {
			return Todo{}, err
		}
	}
	return t, nil
}

func InsertTodo(title, text string) (int, error) {
	query := "INSERT INTO todos (title, text) VALUES (?, ?)"

	result, err := db.DB.Exec(query, title, text)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func UpdateTodo(id int) error {
	query := "UPDATE todos SET completed = 1 WHERE todoId = ?"

	_, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
