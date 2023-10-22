package handlers_test

import (
	"fmt"
	"galaxyed/nginx-be/internal/handlers"
	"galaxyed/nginx-be/internal/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestParseUpstream(t *testing.T) {
// 	// Create a test file
// 	file, err := os.CreateTemp("", "nginx.conf")
// 	assert.NoError(t, err)
// 	defer os.Remove(file.Name())

// 	// Write the test configuration to the file
// 	_, err = file.WriteString(`# This is a comment

// upstream example.com {
//     server example.com:80;
//     server example.com:81;
// }`)

// 	// Read the contents of the file
// 	content, err := os.ReadFile(file.Name())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Parse the test configuration file
// 	upstreams, err := handlers.ParseUpstreams(string(content))

// 	// Log errors
// 	if err != nil {
// 		t.Log(err)
// 	}

// 	// Check that the parsed models match the expected values
// 	expectedUpstreams := []models.NginxUpstream{
// 		{
// 			UpstreamName:      "example.com",
// 			UpstreamServers:   []string{"example.com:80", "example.com:81"},
// 			UpstreamKeepalive: "",
// 		},
// 	}
// 	assert.Equal(t, expectedUpstreams, upstreams)
// }

// func TestParseServer(t *testing.T) {
// 	config := `
//         server {
//             listen 80;
//             server_name example.com;

// 			include /etc/nginx/includes/*.conf;

// 			proxy_hide_header X-Powered-By;
// 			proxy_pass_header Server;
// 			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

//             location / {
//                 proxy_pass http://backend;
//             }
//         }
//     `
// 	expectedServers := []models.NginxServer{
// 		{
// 			ServerName: "example.com",
// 			ServerPort: "80",
// 			Locations: []models.NginxLocation{
// 				{
// 					LocationPath:      "/",
// 					LocationProxyPass: "http://backend",
// 				},
// 			},
// 			Headers: []models.NginxHeader{
// 				{
// 					HeaderAction: "proxy_hide_header",
// 					HeaderName:   "X-Powered-By",
// 					HeaderValue:  "",
// 				},
// 				{
// 					HeaderAction: "proxy_pass_header",
// 					HeaderName:   "Server",
// 					HeaderValue:  "",
// 				},
// 				{
// 					HeaderAction: "proxy_set_header",
// 					HeaderName:   "X-Forwarded-For",
// 					HeaderValue:  "$proxy_add_x_forwarded_for",
// 				},
// 			},
// 			ProxyProps: models.NginxProxyProps{
// 				HideHeaders: []string{"X-Powered-By"},
// 				PassHeaders: []string{"Server"},
// 				SetHeaders: []models.SetHeaders{
// 					{
// 						Header: "X-Forwarded-For",
// 						Value:  "$proxy_add_x_forwarded_for",
// 					},
// 				},
// 			},
// 			Includes: []string{"/etc/nginx/includes/*.conf"},
// 		},
// 	}

// 	server, err := handlers.ParseServers(config)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedServers, server)

// }

// Test ParseServersAndUpstreams
func TestParseServersAndUpstreams(t *testing.T) {
	// Create a test file
	file, err := os.CreateTemp("", "nginx.conf")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	// Write the test configuration to the file
	_, err = file.WriteString(`
# This is a comment

upstream example.com {
	server example.com:80;
	server example.com:81;
}

server {
	listen 80;
	server_name example.com;

	include /etc/nginx/includes/*.conf;

	proxy_hide_header X-Powered-By;
	proxy_pass_header Server;

	location / {
		proxy_hide_header X-Powered-By;
		proxy_pass http://backend;
	}
}
`)

	// Read the contents of the file
	content, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the test configuration file
	servers, upstreams, err := handlers.ParseServersAndUpstreams(string(content))

	// fmt print json pretty
	fmt.Printf("%+v\n", servers)

	// Log errors
	if err != nil {
		t.Log(err)
	}

	// Check that the parsed models match the expected values
	expectedServers := []models.NginxServer{
		{
			ServerName: "example.coms",
			ServerPort: "80",
			Locations: []models.NginxLocation{
				{
					LocationPath:      "/",
					LocationProxyPass: "http://backend",
					ProxyProps: models.NginxProxyProps{
						HideHeaders: []string{"X-Powered-By"},
					},
				},
			},
			ProxyProps: models.NginxProxyProps{
				HideHeaders: []string{"X-Powered-By"},
				PassHeaders: []string{"Server"},
			},
			Includes: []string{"/etc/nginx/includes/*.conf"},
		},
	}
	expectedUpstreams := []models.NginxUpstream{
		{
			UpstreamName: "example.com",

			UpstreamServers:   []string{"example.com:80", "example.com:81"},
			UpstreamKeepalive: "",
		},
	}
	assert.Equal(t, expectedServers, servers)
	assert.Equal(t, expectedUpstreams, upstreams)
}
