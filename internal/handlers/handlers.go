package handlers

import (
	"encoding/json"
	"galaxyed/nginx-be/internal/services"
	"net/http"
)

type Handlers struct {
	nginxServices services.NginxServices
}

func NewHandlers(nginxServices services.NginxServices) *Handlers {
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

	json.NewEncoder(w).Encode(nginxHttp)
}
