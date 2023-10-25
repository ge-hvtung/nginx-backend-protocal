package httpproxy_test

import (
	"fmt"
	"testing"

	"github.com/tufanbarisyildirim/gonginx"
	httpproxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

func TestHttpProxy(t *testing.T) {
	// Create new HttpProxy
	httpProxy := httpproxy.HttpProxy{}

	// Add proxy prop
	httpProxy.AddProp(httpproxy.ProxyHideHeader, "X-Frame-Options")
	httpProxy.AddProp(httpproxy.ProxyHideHeader, "X-Content-Type-Options")
	httpProxy.AddProp(httpproxy.ProxyHideHeader, "Content-Security-Policy")

	// Proxy Add Header
	httpProxy.AddProp(httpproxy.ProxySetHeader, "X-Real-IP $remote_addr")
	httpProxy.AddProp(httpproxy.ProxySetHeader, "X-Forwarded-For $proxy_add_x_forwarded_for")

	// Dump to Directive
	proxyPropDirective := httpproxy.ProxyPropDirective{}

	proxyPropDirective.Dump(&httpProxy)

	for _, directive := range proxyPropDirective.Directives {
		fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))
	}
}
