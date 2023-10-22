package config

import (
	"os"
)

func GetNginxLocation() string {
	// Get the location of the Nginx configuration directory from the NGINX_CONF_PATH environment variable
	location := os.Getenv("NGINX_CONF_PATH")
	if location == "" {
		location = "/etc/nginx"
	}
	return location
}
