package handlers

import (
	"encoding/json"
	"fmt"
	"galaxyed/nginx-be/internal/models"
	"strings"
)

func ParseServersAndUpstreams(config string) ([]models.NginxServer, []models.NginxUpstream, error) {
	var servers []models.NginxServer
	var currentServer *models.NginxServer
	var currentLocation *models.NginxLocation
	var upstreams []models.NginxUpstream
	var currentUpstream *models.NginxUpstream

	for _, line := range strings.Split(config, "\n") {
		line = strings.TrimSpace(line)

		switch {
		case line == "":
			continue
		case strings.HasPrefix(line, "#"):
			continue
		case strings.HasPrefix(line, "server {"):
			currentServer = &models.NginxServer{}
			currentLocation = nil
		case strings.HasPrefix(line, "location "):
			locationPath := strings.TrimSuffix(strings.TrimPrefix(line, "location "), " {")
			currentLocation = &models.NginxLocation{LocationPath: locationPath}
		case strings.HasPrefix(line, "server_name "):
			currentServer.ServerName = strings.TrimSuffix(strings.TrimPrefix(line, "server_name "), ";")
		case strings.HasPrefix(line, "listen "):
			currentServer.ServerPort = strings.TrimSuffix(strings.TrimPrefix(line, "listen "), ";")
		case strings.HasPrefix(line, "proxy_pass "):
			currentLocation.LocationProxyPass = strings.TrimSuffix(strings.TrimPrefix(line, "proxy_pass "), ";")
		case strings.HasPrefix(line, "proxy_hide_header "):
			headerName := strings.TrimSuffix(strings.TrimPrefix(line, "proxy_hide_header "), ";")
			// Add to current location if it exists, otherwise add to current server
			if currentLocation != nil {
				currentLocation.ProxyProps.HideHeaders = append(currentLocation.ProxyProps.HideHeaders, headerName)
			} else {
				currentServer.ProxyProps.HideHeaders = append(currentServer.ProxyProps.HideHeaders, headerName)
			}
		case strings.HasPrefix(line, "proxy_pass_header "):
			headerName := strings.TrimSuffix(strings.TrimPrefix(line, "proxy_pass_header "), ";")
			// Add to current location if it exists, otherwise add to current server
			if currentLocation != nil {
				currentLocation.ProxyProps.PassHeaders = append(currentLocation.ProxyProps.PassHeaders, headerName)
			} else {
				currentServer.ProxyProps.PassHeaders = append(currentServer.ProxyProps.PassHeaders, headerName)
			}
		case strings.HasPrefix(line, "proxy_set_header "):
			headerParts := strings.Split(strings.TrimSuffix(strings.TrimPrefix(line, "proxy_set_header "), ";"), " ")
			// Add to current location if it exists, otherwise add to current server
			if currentLocation != nil {
				currentLocation.ProxyProps.SetHeaders = append(currentLocation.ProxyProps.SetHeaders, models.SetHeaders{Header: headerParts[0], Value: strings.Join(headerParts[1:], " ")})
			} else {
				currentServer.ProxyProps.SetHeaders = append(currentServer.ProxyProps.SetHeaders, models.SetHeaders{Header: headerParts[0], Value: strings.Join(headerParts[1:], " ")})
			}
		case strings.HasPrefix(line, "include "):
			currentServer.Includes = append(currentServer.Includes, strings.TrimSuffix(strings.TrimPrefix(line, "include "), ";"))
		case line == "}":
			if currentServer != nil {
				if currentLocation != nil {
					currentServer.Locations = append(currentServer.Locations, *currentLocation)
					currentLocation = nil
				}
				servers = append(servers, *currentServer)
				currentServer = nil
			}
		case strings.HasPrefix(line, "upstream "):
			upstreamName := strings.TrimSuffix(strings.TrimPrefix(line, "upstream "), " {")
			currentUpstream = &models.NginxUpstream{
				UpstreamName: upstreamName,
			}
			if upstreams == nil {
				upstreams = []models.NginxUpstream{}
			}
			upstreams = append(upstreams, *currentUpstream)
		case currentUpstream != nil && strings.HasPrefix(line, "server "):
			server := strings.TrimSuffix(strings.TrimPrefix(line, "server "), ";")

			if currentUpstream.UpstreamServers == nil {
				currentUpstream.UpstreamServers = []string{}
			}

			currentUpstream.UpstreamServers = append(currentUpstream.UpstreamServers, server)

			for i, upstream := range upstreams {
				if upstream.UpstreamName == currentUpstream.UpstreamName {
					upstreams[i] = *currentUpstream
				}
			}
		case line == "}":
			if currentUpstream != nil {
				currentUpstream = nil
			}
		default:
			return nil, nil, fmt.Errorf("unknown line: %s", line)
		}
	}

	// fmt print json pretty
	jsonBytes, err := json.MarshalIndent(servers, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, nil, err
	}
	fmt.Println(string(jsonBytes))

	return servers, upstreams, nil
}
