package httpupstream_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tufanbarisyildirim/gonginx/parser"
	httpupstream "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_upstream"
)

func TestParseUpstreamBlock(t *testing.T) {
	// Test data
	data := parser.NewStringParser(`http {
		upstream backend {
			keepalive 32;
			hash $request_uri;
			least_conn;
			ip_hash;
			server 127.0.0.1:8080;
			server localhost:8080;
		}
	}`)
	conf := data.Parse()
	upstreams := conf.FindUpstreams()
	upstream, err := httpupstream.ParseUpstreamBlock(upstreams[0])

	fmt.Println("upstream", upstream, err)
	assert.Nil(t, err)

	// Test upstream
	assert.Equal(t, "backend", upstreams[0].UpstreamName)
	assert.Equal(t, "32", upstream.Keepalive)
	assert.Equal(t, true, upstream.LeastConn)
	assert.Equal(t, true, upstream.IpHash)
}
