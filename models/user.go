package models

import (
	"a.com/http/db"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID   int
	Username string
	Password string
}

func InsertUser(username, password string) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	query := "INSERT INTO users (username,password) VALUES (?, ?)"

	_, err = db.DB.Exec(query, username, string(hashedPass))
	if err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(username, password string) error {

	var hashedPass []byte

	query := "SELECT password FROM users WHERE username = ?"

	err := db.DB.QueryRow(query, username).Scan(&hashedPass)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return err
	}
	return nil
}
