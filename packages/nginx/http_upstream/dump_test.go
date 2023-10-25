package httpupstream_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tufanbarisyildirim/gonginx"
	upstream "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_upstream"
)

func TestUpstreamContext(t *testing.T) {
	http_upstream := upstream.Upstream{
		UpstreamName: "backend",
		Servers: []upstream.Server{
			{
				Host:            "127.0.0.1",
				Port:            "8080",
				Weight:          1,
				MaxConns:        0,
				MaxFails:        0,
				MaxConnsTimeout: "0s",
			},
		},
		Keepalive: "32",
		Hash:      "ip_hash",
		LeastConn: true,
		IpHash:    true,
	}

	upstreamContext := upstream.UpstreamContext{
		Upstream: http_upstream,
	}

	directive := upstream.UpstreamDirective{}
	directive.Dump(&upstreamContext)

	expected := `upstream backend {
    keepalive 32;
    hash ip_hash;
    least_conn;
    ip_hash;
    server 127.0.0.1:8080 weight=1 max_conns_timeout=0s;
}`

	actual := gonginx.DumpDirective(directive.Dump(&upstreamContext), gonginx.IndentedStyle)

	assert.Equal(t, expected, actual)
}
