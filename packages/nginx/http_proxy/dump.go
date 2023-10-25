package httpproxy

import "github.com/tufanbarisyildirim/gonginx"

func (c *HttpProxy) Init() {
	c.Props = make([]ProxyProp, 0)
}

// Add proxy prop
func (c *HttpProxy) AddProp(name, value string) {
	c.Props = append(c.Props, ProxyProp{
		Name:  name,
		Value: value,
	})
}

type ProxyPropDirective struct {
	Directives []gonginx.IDirective
}

// Dump to Directive
func (d *ProxyPropDirective) Dump(p *HttpProxy) {
	for _, prop := range p.Props {
		directive := gonginx.Directive{
			Name:       prop.Name,
			Parameters: []string{prop.Value},
		}

		d.Directives = append(d.Directives, &directive)
	}
}
