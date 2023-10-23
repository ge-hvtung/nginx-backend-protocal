package models

// New NginxHttp struct
type NgxHttp struct {
	Upstreams []NgxUpstream `json:"upstreams"`
	Servers   []NgxServer   `json:"servers"`
}

type NgxUpstream struct {
	UpstreamName      string              `json:"upstream_name"`
	UpstreamServers   []NgxUpstreamServer `json:"upstream_servers"`
	UpstreamKeepalive string              `json:"upstream_keepalive"`
}

type NgxUpstreamServer struct {
	Address string `json:"address"`
}

type NgxServer struct {
	ServerName []string      `json:"server_name"`
	ServerPort string        `json:"server_port"`
	Listen     []string      `json:"listen"`
	ProxyProps []string      `json:"proxy_props"`
	Locations  []NgxLocation `json:"locations"`
	Includes   []string      `json:"includes"`
}

type NgxLocation struct {
	LocationPath      string   `json:"location_path"`
	LocationProxyPass string   `json:"location_proxy_pass"`
	ProxyProps        []string `json:"proxy_props"`
}
