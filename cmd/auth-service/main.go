package main

import (
	"auth-service/internal/app"
	"auth-service/internal/routes"
	"log"
	"net/http"
)

func main() {
	app.Initialize()

	mux := http.NewServeMux()

	routes.Configure(mux)

	address := ":8080"
	log.Printf("Server running on http://localhost%s/", address)
	log.Fatal(http.ListenAndServe(address, mux))
}
