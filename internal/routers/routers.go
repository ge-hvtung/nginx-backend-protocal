package routers

import (
	"net/http"

	"github.com/tunghauvan/nginx-backend-protocal/internal/handlers"

	"github.com/gorilla/mux"
	// httpSwagger "github.com/swaggo/http-swagger"
)

// SetRoutes sets up the routes for the application
func SetRoutes(r *mux.Router, h *handlers.Handlers) {
	// /api/v1 routes
	api := r.PathPrefix("/api/v1").Subrouter()

	// /api/v1/nginx routes
	nginx := api.PathPrefix("/nginx").Subrouter()
	nginx.HandleFunc("/files", h.GetNginxFiles).Methods(http.MethodGet)

}
