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
	r.HandleFunc("/configs", h.GetNgxConfig).Methods(http.MethodGet)

	// Get All Upstreams
	r.HandleFunc("/upstreams", h.GetUpstreams).Methods(http.MethodGet)

	// Get All Servers
	r.HandleFunc("/servers", h.GetServers).Methods(http.MethodGet)
}
