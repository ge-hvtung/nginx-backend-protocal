package httpupstream

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tufanbarisyildirim/gonginx"
)

// Parse Upstream Block
func ParseUpstreamBlock(block gonginx.IBlock) (upstream Upstream, err error) {
	// Print upstream block
	fmt.Println(gonginx.DumpBlock(block, gonginx.IndentedStyle))

	// Parse upstream block
	upstream = Upstream{}
	upstream.UpstreamName = ""
	upstream.Keepalive = ""
	upstream.Hash = ""
	upstream.LeastConn = false
	upstream.IpHash = false
	upstream.Servers = []Server{}

	// Parse directives
	for _, directive := range block.GetDirectives() {
		// Parse upstream_name
		if directive.GetName() == "upstream" {
			upstream.UpstreamName = directive.GetParameters()[0]
			continue
		}

		// Parse keepalive
		if directive.GetName() == "keepalive" {
			upstream.Keepalive = directive.GetParameters()[0]
			continue
		}

		// Parse least_conn
		if directive.GetName() == "least_conn" {
			upstream.LeastConn = true
			continue
		}

		// Parse ip_hash
		if directive.GetName() == "ip_hash" {
			upstream.IpHash = true
			continue
		}

		// Parse server
		if directive.GetName() == "server" {
			server := Server{}
			server.Host = ""
			server.Port = ""
			server.Weight = 0
			server.MaxConns = 0
			server.MaxFails = 0
			server.MaxConnsTimeout = ""

			// Get parameters
			parameters := directive.GetParameters()

			// Parse host and port
			host_port := strings.Split(parameters[0], ":")
			server.Host = host_port[0]
			server.Port = host_port[1]

			// Parse weight
			if len(parameters) > 1 {
				weight := strings.Split(parameters[1], "=")
				server.Weight, _ = stringToInt(weight[1])
			}

			// Parse max_conns parameter starts with max_conns_timeout
			if len(parameters) > 2 {
				max_conns := strings.Split(parameters[2], "=")
				server.MaxConns, _ = stringToInt(max_conns[1])
			}

			// Parse max_fails
			if len(parameters) > 3 {
				max_fails := strings.Split(parameters[3], "=")
				server.MaxFails, _ = stringToInt(max_fails[1])
			}

			// Parse max_conns_timeout
			if len(parameters) > 4 {
				server.MaxConnsTimeout = parameters[4]
			}

			// Append server to upstream
			upstream.Servers = append(upstream.Servers, server)
			continue
		}
	}

	return upstream, nil
}

func stringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
