package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tunghauvan/nginx-backend-protocal/internal/config"
	"github.com/tunghauvan/nginx-backend-protocal/internal/handlers"
	"github.com/tunghauvan/nginx-backend-protocal/internal/routers"
	"github.com/tunghauvan/nginx-backend-protocal/internal/services"

	"github.com/gorilla/mux"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <nginx-config-file> <nginx-config-directory>")
	}

	// Get the arguments
	nginxConfigFile := os.Args[1]
	nginxConfigDir := os.Args[2]

	// Load configuration
	config.GetNginxLocation()

	// Initialize router
	router := mux.NewRouter()

	// Initialize services
	ngxSvc := services.NewNgxService()
	ngxSvc.SetConfig(nginxConfigFile)
	ngxSvc.SetDirectory(nginxConfigDir)
	ngxSvc.ReadNginxConfiguration()

	// Initialize handlers
	nginxHandler := handlers.NewHandlers(ngxSvc)

	// Set up routes
	routers.SetRoutes(router, nginxHandler)

	// Start server
	log.Printf("Server listening on port %s", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
