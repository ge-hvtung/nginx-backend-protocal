package models

type NginxHttp struct {
	Servers  []NginxServer   `json:"servers"`
	Upstrems []NginxUpstream `json:"upstreams"`
}

type NginxUpstream struct {
	UpstreamName      string   `json:"upstream_name"`
	UpstreamServers   []string `json:"upstream_servers"`
	UpstreamKeepalive string   `json:"upstream_keepalive"`
}

type NginxLocation struct {
	LocationPath      string          `json:"location_path"`
	LocationProxyPass string          `json:"location_proxy_pass"`
	ProxyProps        NginxProxyProps `json:"proxy_props"`
}

type NginxHeader struct {
	HeaderAction string `json:"header_action"`
	HeaderName   string `json:"header_name"`
	HeaderValue  string `json:"header_value"`
}

type NginxProxyProps struct {
	HideHeaders []string     `json:"proxy_hide_headers"`
	PassHeaders []string     `json:"proxy_pass_headers"`
	SetHeaders  []SetHeaders `json:"proxy_set_headers"`
}

type SetHeaders struct {
	Header string `json:"header"`
	Value  string `json:"value"`
}

type NginxServer struct {
	ServerName string          `json:"server_name"`
	ServerPort string          `json:"server_port"`
	Locations  []NginxLocation `json:"locations"`
	Includes   []string        `json:"includes"`
	ProxyProps NginxProxyProps `json:"proxy_props"`
}

type NginxService struct {
	Upstreams []NginxUpstream `json:"upstreams"`
	Servers   []NginxServer   `json:"servers"`
}

type NginxUpstreamRepository interface {
	GetAll() ([]NginxUpstream, error)
	SetAll(upstreams []NginxUpstream) error
}

type NginxServerRepository interface {
	GetAll() ([]NginxServer, error)
	SetAll(servers []NginxServer) error

	GetServerByName(name string) (NginxServer, error)
}
