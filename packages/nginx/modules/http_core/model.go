package httpcore

// define client_body_buffer_size object
// ClientBodyBufferSize directive
// Syntax:	client_body_buffer_size size
// Default:
// Context:	http, server, location
//
// Sets buffer size for reading client request body. In case the request body is larger than the buffer, the whole body or only its part is written to a temporary file.
// By default, buffer size is equal to two memory pages. This is 8K on x86, other 32-bit platforms, and x86-64. It is usually 16K on other 64-bit platforms.
// The first line sets buffer size to 1K bytes. The second line sets the buffer size to 10 bytes, which means that nginx will write only the first part of the request body into a temporary file.
// client_body_buffer_size 1k;
// client_body_buffer_size 10;
type ClientBodyBufferSize struct {
	// value field
	Value string `json:"value"`
}
