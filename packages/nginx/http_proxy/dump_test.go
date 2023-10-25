package proxy_test

import (
	"fmt"
	"testing"

	"github.com/tufanbarisyildirim/gonginx"
	proxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

func TestHttpProxy(t *testing.T) {
	// Create new HttpProxy
	httpProxy := proxy.HttpProxy{}

	// Add proxy prop
	httpProxy.AddProp(proxy.ProxyHideHeader, "X-Frame-Options")
	httpProxy.AddProp(proxy.ProxyHideHeader, "X-Content-Type-Options")
	httpProxy.AddProp(proxy.ProxyHideHeader, "Content-Security-Policy")

	// Proxy Add Header
	httpProxy.AddProp(proxy.ProxySetHeader, "X-Real-IP $remote_addr")
	httpProxy.AddProp(proxy.ProxySetHeader, "X-Forwarded-For $proxy_add_x_forwarded_for")

	// Dump to Directive
	proxyPropDirective := proxy.ProxyPropDirective{}

	proxyPropDirective.Dump(&httpProxy)

	for _, directive := range proxyPropDirective.Directives {
		fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))
	}
}
