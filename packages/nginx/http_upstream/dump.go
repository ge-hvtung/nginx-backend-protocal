package httpupstream

import (
	"strconv"

	"github.com/tufanbarisyildirim/gonginx"
)

type UpstreamDirective struct {
	Directives []gonginx.IDirective
}

// Dump to Directive
func (d *UpstreamDirective) Dump(p *UpstreamContext) gonginx.IDirective {
	directives := UpstreamDirective{
		Directives: make([]gonginx.IDirective, 0),
	}

	if p.Upstream.Keepalive != "" {
		keepaliveDirective := gonginx.Directive{
			Name:       "keepalive",
			Parameters: []string{p.Upstream.Keepalive},
		}

		directives.Add(&keepaliveDirective)
	}

	if p.Upstream.Hash != "" {
		hashDirective := gonginx.Directive{
			Name:       "hash",
			Parameters: []string{p.Upstream.Hash},
		}

		directives.Add(&hashDirective)
	}

	if p.Upstream.LeastConn {
		leastConnDirective := gonginx.Directive{
			Name: "least_conn",
		}

		directives.Add(&leastConnDirective)
	}

	if p.Upstream.IpHash {
		ipHashDirective := gonginx.Directive{
			Name: "ip_hash",
		}

		directives.Add(&ipHashDirective)
	}

	for _, server := range p.Upstream.Servers {
		serverDirective := gonginx.Directive{
			Name:       "server",
			Parameters: []string{server.Host + ":" + server.Port},
		}

		if server.Weight != 0 {
			serverDirective.Parameters = append(serverDirective.Parameters, "weight="+toString(server.Weight))
		}

		if server.MaxConns != 0 {
			serverDirective.Parameters = append(serverDirective.Parameters, "max_conns="+toString(server.MaxConns))
		}

		if server.MaxFails != 0 {
			serverDirective.Parameters = append(serverDirective.Parameters, "max_fails="+toString(server.MaxFails))
		}

		if server.MaxConnsTimeout != "" {
			serverDirective.Parameters = append(serverDirective.Parameters, "max_conns_timeout="+server.MaxConnsTimeout)
		}

		directives.Add(&serverDirective)
	}

	upstream_directive := gonginx.Directive{
		Name:       "upstream",
		Parameters: []string{p.Upstream.UpstreamName},
		Block: &gonginx.Block{
			Directives: directives.Directives,
		},
	}
	return &upstream_directive
}

// directive add
func (d *UpstreamDirective) Add(directive gonginx.IDirective) {
	d.Directives = append(d.Directives, directive)
}

func toString(i int) string {
	return strconv.Itoa(i)
}
