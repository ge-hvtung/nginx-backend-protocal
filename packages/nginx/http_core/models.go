package core

// Ref: http://nginx.org/en/docs/http/ngx_http_core_module.html

import (
	access "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_access"
	proxy "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_proxy"
)

type HttpContext struct {
	ClientPropsHttp
	access.HttpAccessContext

	ErrorPageContext

	// Proxy
	Proxy []*proxy.HttpProxy `json:"proxy"`

	Servers []*ServerContext `json:"server"` // ServerContext
}

type ServerContext struct {
	ServerNames []string        `json:"server_names"`
	Listens     []ListenContext `json:"listens"` // ListenContext

	ClientPropsServer
	access.HttpAccessContext

	// Properties
	CoreProps ClientPropsServer `json:"core_props"`
	Proxy     proxy.HttpProxy   `json:"proxy"`

	// ErrorPageContext
	ErrorPageContext []ErrorPageContext `json:"error_page"`

	// Paths
	Locations []*LocationContext `json:"location"`
}

type ListenContext struct {
	// Listen can be a number, IP:port, or a path
	Listen []string `json:"listen"`

	// SSL
	SSL bool `json:"ssl"`

	// Port
	Ports []string `json:"ports"`
}

type ClientPropsServer struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}

type LocationContext struct {
	CoreProps ClientPropsLocation `json:"core_props"`
	access.HttpAccessContext

	// Paths
	Paths []string `json:"paths"`

	// Error page
	ErrorPageContext []ErrorPageContext `json:"error_page"`
}

// define error_page
// Syntax: error_page code ... [=[response]] uri;
type ErrorPageContext struct {
	// code can be a number or "*"
	Codes []int `json:"code"`

	// Return code if any
	Response string `json:"response"`

	// uri can be text, variable, or a URI
	URI string `json:"uri"`
}

type ClientPropsHttp struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}

type ClientPropsLocation struct {
	AbsoluteRedirect        string `json:"absolute_redirect"`
	Aio                     string `json:"aio"`
	AioWrite                string `json:"aio_write"`
	ChunkedTransferEncoding string `json:"chunked_transfer_encoding"`
	ClientBodyBufferSize    string `json:"client_body_buffer_size"`
	KeepaliveRequests       string `json:"keepalive_requests"`
	ClientMaxBodySize       string `json:"client_max_body_size"`
}
