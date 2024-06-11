package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"GeminiChallenge/routes"
)

func main() {
	mux := mux.NewRouter()
	port := 8080

	mux.HandleFunc("/signup", routes.EssayRoute)

	// CORS middleware
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	// Start the server on port 8080
	log.Printf("Server listening on port %d", port)
	err := http.ListenAndServe(":8080", handlers.CORS(headers, origins, methods)(mux))
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}