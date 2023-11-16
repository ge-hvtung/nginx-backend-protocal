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
	ClientPropsServer
	access.HttpAccessContext
	ServerNames      []string           `json:"server_names"`
	Listens          []ListenContext    `json:"listens"` // ListenContext
	CoreProps        ClientPropsServer  `json:"core_props"`
	Proxy            proxy.HttpProxy    `json:"proxy"`
	ErrorPageContext []ErrorPageContext `json:"error_page"`
	Locations        []LocationContext  `json:"location"`
}

type ListenContext struct {
	Listen []string `json:"listen"`
	SSL    bool     `json:"ssl"`
	Ports  []string `json:"ports"`
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
	Path       string             `json:"path"`
	ErrorPages []ErrorPageContext `json:"error_page"`
	Proxy      proxy.HttpProxy    `json:"proxy"`
}

// define error_page
// Syntax: error_page code ... [=[response]] uri;
type ErrorPageContext struct {
	Codes    []int  `json:"code"`
	Response string `json:"response"`
	URI      string `json:"uri"`
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
