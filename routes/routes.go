package routes

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /todos", getTodos)
	mux.HandleFunc("GET /todos/{id}", getTodo)
	mux.HandleFunc("POST /todos/create", postTodo)
	mux.HandleFunc("PATCH /todos/update/{id}", patchTodo)

	mux.HandleFunc("POST /users/create", postUser)
	mux.HandleFunc("POST /users/login", loginUserPost)
}
