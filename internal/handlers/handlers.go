package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tunghauvan/nginx-backend-protocal/internal/services"
)

type Handlers struct {
	nginxServices *services.NgxService
}

// NewHandlers returns a new Handlers struct
func NewHandlers(nginxServices *services.NgxService) *Handlers {
	return &Handlers{
		nginxServices: nginxServices,
	}
}

// Handler to get the Nginx HTTP configurations
func (h *Handlers) GetNginxHttp(w http.ResponseWriter, r *http.Request) {
	nginxHttp, err := h.nginxServices.GetNginxHttp()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	// Write response
	json.NewEncoder(w).Encode(nginxHttp)
}

func (h *Handlers) GetNgxConfig(w http.ResponseWriter, r *http.Request) {
	h.nginxServices.ParseNginxConfiguration()

	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	// Write response
	json.NewEncoder(w).Encode(nil)
}

// Handler to get all upstreams
func (h *Handlers) GetUpstreams(w http.ResponseWriter, r *http.Request) {
	upstreams, err := h.nginxServices.GetUpstreams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	// Write response
	json.NewEncoder(w).Encode(upstreams)
}

// Handler to get all servers
func (h *Handlers) GetServers(w http.ResponseWriter, r *http.Request) {
	servers, err := h.nginxServices.GetServers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	// Write response
	json.NewEncoder(w).Encode(servers)
}
