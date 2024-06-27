package routes

import (
	"net/http"

	"a.com/http/models"
)

func postUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == "" {
		http.Error(w, "cannot be blank", http.StatusInternalServerError)
		return
	}
	if password == "" {
		http.Error(w, "cannot be blank", http.StatusInternalServerError)
		return
	}

	err = models.InsertUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func loginUserPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == "" {
		http.Error(w, "cannot be blank", http.StatusInternalServerError)
		return
	}
	if password == "" {
		http.Error(w, "cannot be blank", http.StatusInternalServerError)
		return
	}

	err = models.AuthenticateUser(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusSeeOther)
}
