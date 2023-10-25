package httpupstream

// Ref: http://nginx.org/en/docs/http/ngx_http_upstream_module.html

// Upstream
type Upstream struct {
	UpstreamName string   `json:"upstream_name"`
	Servers      []Server `json:"servers"`
	Keepalive    string   `json:"keepalive"`
	Hash         string   `json:"hash"`
	LeastConn    bool     `json:"least_conn"`
	IpHash       bool     `json:"ip_hash"`
}

// Server
type Server struct {
	Host            string `json:"host"`
	Port            string `json:"port"`
	Weight          int    `json:"weight"`
	MaxConns        int    `json:"max_conns"`
	MaxFails        int    `json:"max_fails"`
	MaxConnsTimeout string `json:"max_conns_timeout"`
}

// UpstreamContext
type UpstreamContext struct {
	Upstream Upstream
}

// UpstreamContexts
type UpstreamContexts []UpstreamContext
