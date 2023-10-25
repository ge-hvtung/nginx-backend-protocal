package proxy

// Make definition for Nginx HTTP Proxy module
// Ref: https://nginx.org/en/docs/http/ngx_http_proxy_module.html

type HttpProxy struct {
	Props     []ProxyProp `json:"props"`
	ProxyPass string      `json:"proxy_pass"`
}

type ProxyProp struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Alloweed Prop names
const (
	ProxySetHeader   = "proxy_set_header"
	ProxyHideHeader  = "proxy_hide_header"
	ProxyPassHeader  = "proxy_pass_header"
	ProxyPassBody    = "proxy_pass_body"
	ProxyPass        = "proxy_pass"
	ProxyMethod      = "proxy_method"
	ProxyHttpVersion = "proxy_http_version"
	ProxySetBody     = "proxy_set_body"
)
