package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tunghauvan/nginx-backend-protocal/internal/services"
)

type Handlers struct {
	nginxServices *services.NginxService
}

// NewHandlers returns a new Handlers struct
func NewHandlers(nginxServices *services.NginxService) *Handlers {
	return &Handlers{
		nginxServices: nginxServices,
	}
}

// Handler to get the List of Nginx files
func (h *Handlers) GetNginxFiles(w http.ResponseWriter, r *http.Request) {
	// Parse the "name" query parameter from the request URL
	name := r.URL.Query().Get("name")
	format := r.URL.Query().Get("format")

	// Format Json response
	w.Header().Add("Content-Type", "application/json")

	if format == "" {
		format = "json"
	} 


	if name != "" {
		contents, err := h.nginxServices.GetNginxFile(name, format)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(contents)

	} else {
		contents, err := h.nginxServices.GetNginxFiles()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(contents)

	}

	// Write response
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
