package core

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unicode"

	"github.com/tufanbarisyildirim/gonginx"

	proxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

// Translate location context to nginx config using gonginx
func (c *LocationContext) ToNginx() string {
	// Init Empty Directives
	directives := &Directives{}

	// Add http_access to directives
	directives.AddHttpAccessContext(c.HttpAccessContext)

	// Add CoreProps to directives
	directives.AddCoreProps(reflect.ValueOf(c.CoreProps))

	// Add error_page to directives
	directives.AddErrorPageContext(c.ErrorPages)

	location_directive := gonginx.Directive{
		Name:       "location",
		Parameters: []string{c.Path},
		Block: &gonginx.Block{
			Directives: directives.Directives,
		},
	}

	return gonginx.DumpDirective(&location_directive, gonginx.IndentedStyle)
}

// Dump server context to nginx config
func (c *ServerContext) ToNginx() gonginx.IDirective {
	// Init Empty Directives
	directives := &Directives{}

	// Add server_name directive
	if len(c.ServerNames) > 0 {
		server_name_directive := gonginx.Directive{
			Name:       "server_name",
			Parameters: c.ServerNames,
		}

		directives.AddDirective(&server_name_directive)
	}

	// Add listen directive multiple times
	if len(c.Listens) > 0 {
		for _, listen := range c.Listens {
			listen_directive := gonginx.Directive{
				Name:       "listen",
				Parameters: listen.Listen,
			}

			directives.AddDirective(&listen_directive)
		}
	}

	// Add http_access to directives
	directives.AddHttpAccessContext(c.HttpAccessContext)

	// Add CoreProps to directives
	directives.AddCoreProps(reflect.ValueOf(c.CoreProps))

	// Dump Proxy to Directive
	proxyPropDirective := proxy.ProxyPropDirective{}
	proxyPropDirective.Dump(&c.Proxy)
	for _, directive := range proxyPropDirective.Directives {
		directives.AddDirective(directive)
	}

	// Add error_page to directives
	directives.AddErrorPageContext(c.ErrorPageContext)

	server_directive := gonginx.Directive{
		Name: "server",
		Block: &gonginx.Block{
			Directives: directives.Directives,
		},
	}

	// Return server directive as IDirective
	return &server_directive
}

func intSliceToString(intSlice []int) string {
	var result string
	for _, i := range intSlice {
		result += fmt.Sprintf("%d ", i)
	}
	return result
}

func toLowerSnakeCase(s string) string {
	var result string
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i > 0 {
				result += "_"
			}
			result += string(unicode.ToLower(c))
		} else {
			result += string(c)
		}
	}
	return result
}

// Dump location context to json string
func (c *LocationContext) Dump() string {
	// Convert the LocationContext object to a JSON string
	jsonString, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Print the resulting JSON string
	fmt.Println(string(jsonString))
	return string(jsonString)
}
