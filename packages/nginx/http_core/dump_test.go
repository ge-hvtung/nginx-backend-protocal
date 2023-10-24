package httpcore_test

import (
	"fmt"
	"testing"

	httpaccess "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_access"
	httpcore "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_core"
)

// Test Location toNginx
func TestLocationToNginx(t *testing.T) {
	http_access := &httpaccess.HttpAccessContext{}
	http_access.Allow = append(http_access.Allow, "10.100.2.0/24")

	location_context := &httpcore.LocationContext{
		Location: []string{"/admin"},
	}
	location_context.Allow = append(location_context.Allow, http_access.Allow...)
	location_context.CoreProps.ClientBodyBufferSize = "128k"
	location_context.CoreProps.ClientMaxBodySize = "1m"

	// Error page
	location_context.ErrorPageContext = httpcore.ErrorPageContext{
		Codes:    []int{404, 500},
		URI:      "/50x.html",
		Response: "=200 @error_page_404",
	}

	fmt.Println(location_context.ToNginx())
}

func TestServerToNginx(t *testing.T) {
	server_context := &httpcore.ServerContext{
		ServerNames: []string{"localhost"},
		Listens:     []string{"80", "443"},
	}
	server_context.CoreProps.ClientBodyBufferSize = "128k"
	server_context.CoreProps.ClientMaxBodySize = "1m"

	// Error page
	server_context.ErrorPageContext = httpcore.ErrorPageContext{
		Codes:    []int{404, 500},
		URI:      "/50x.html",
		Response: "=200 @error_page_404",
	}

	fmt.Println(server_context.ToNginx())
}
