package routers

import (
	"net/http"

	"galaxyed/nginx-be/internal/handlers"

	"github.com/gorilla/mux"
)

// SetRoutes sets up the routes for the application
func SetRoutes(r *mux.Router, h *handlers.Handlers) {
	// Get Nginx configuration
	r.HandleFunc("/nginx", h.GetNginxHttp).Methods(http.MethodGet)
}
