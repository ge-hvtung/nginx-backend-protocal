package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"galaxyed/nginx-be/internal/config"
	"galaxyed/nginx-be/internal/handlers"
	"galaxyed/nginx-be/internal/routers"
	"galaxyed/nginx-be/internal/services"
)

func main() {
	// Load configuration
	config.GetNginxLocation()

	// Initialize router
	router := mux.NewRouter()

	// Initialize services
	nginxServices := services.NewNginxServices()

	// Initialize handlers
	nginxHandler := handlers.NewHandlers(nginxServices)

	// Set up routes
	routers.SetRoutes(router, nginxHandler)

	// Start server
	log.Printf("Server listening on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
