package main

import (
	"log"
	"log/slog"
	"net/http"

	"a.com/http/db"
	"a.com/http/routes"
)

func main() {
	db.InitDb()

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	address := ":8080"

	slog.Info("starting server", "address", address)

	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Fatal(err)
	}
}
