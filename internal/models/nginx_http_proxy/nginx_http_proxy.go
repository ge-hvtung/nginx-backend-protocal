package nginxHttpProxy

// Make definition for Nginx HTTP Proxy module
// Ref: https://nginx.org/en/docs/http/ngx_http_proxy_module.html

type NginxHttpProxy struct {
	Props     []NginxHttpProxyProp `json:"props"`
	ProxyPass string               `json:"proxy_pass"`
}

type NginxHttpProxyProp struct {
	PropName  string `json:"prop_name"`
	PropValue string `json:"prop_value"`
}

// Func to validate Nginx HTTP Proxy module
func (n *NginxHttpProxy) Validate() error {
	// TODO: Implement validation
	return nil
}
