package main

import (
	"log"
	"net/http"
	"os"

	// "github.com/tunghauvan/nginx-backend-protocal/internal/config"
	"github.com/tunghauvan/nginx-backend-protocal/internal/handlers"
	"github.com/tunghauvan/nginx-backend-protocal/internal/routers"
	"github.com/tunghauvan/nginx-backend-protocal/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run main.go <nginx-config-directory>")
	}

	// Get the arguments
	nginxConfigDir := os.Args[1]

	// Initialize router
	router := mux.NewRouter()

	// Initialize services
	nginxService := services.NewNginxService()
	nginxService.SetDirectory(nginxConfigDir)
	nginxService.ReadNginxConfiguration()

	// Initialize handlers
	nginxHandler := handlers.NewHandlers(nginxService)

	// Set up routes
	routers.SetRoutes(router, nginxHandler)

	// Start server
	log.Printf("Server listening on port %s", "8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
