package core_test

import (
	"fmt"
	"testing"

	"github.com/tufanbarisyildirim/gonginx"
	access "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_access"
	core "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_core"
	proxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

// Test Location toNginx
func TestLocationToNginx(t *testing.T) {
	http_access := &access.HttpAccessContext{}
	http_access.Allow = append(http_access.Allow, "10.100.2.0/24")

	location_context := &core.LocationContext{
		Location: []string{"/admin"},
	}
	location_context.Allow = append(location_context.Allow, http_access.Allow...)
	location_context.CoreProps.ClientBodyBufferSize = "128k"
	location_context.CoreProps.ClientMaxBodySize = "1m"

	// Error page
	location_context.ErrorPageContext = core.ErrorPageContext{
		Codes:    []int{404, 500},
		URI:      "/50x.html",
		Response: "=200 @error_page_404",
	}

	fmt.Println(location_context.ToNginx())
}

func TestServerToNginx(t *testing.T) {
	server_context := &core.ServerContext{
		ServerNames: []string{"localhost", "example.com"},
		Listens:     []string{"80", "443"},
	}
	server_context.CoreProps.ClientBodyBufferSize = "128k"
	server_context.CoreProps.ClientMaxBodySize = "1m"

	// Remove Proxy Header Access-Control-Allow-Origin
	server_context.Proxy.AddProp(proxy.ProxyHideHeader, "Access-Control-Allow-Origin")
	server_context.Proxy.AddProp(proxy.ProxySetHeader, "Access-Control-Allow-Origin $http_origin")

	// Error page
	server_context.ErrorPageContext = core.ErrorPageContext{
		Codes:    []int{404, 500},
		URI:      "/50x.html",
		Response: "=200 @error_page_404",
	}

	// Create gonginx config

	conf := &gonginx.Config{
		Block: &gonginx.Block{
			Directives: []gonginx.IDirective{
				server_context.ToNginx(),
			},
		},
	}

	fmt.Println(gonginx.DumpConfig(conf, gonginx.IndentedStyle))
}
