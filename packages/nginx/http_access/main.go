package access

// Ref: http://nginx.org/en/docs/http/ngx_http_access_module.html

import "net"

type HttpAccessContext struct {
	Allow []string `json:"allow"`
	Deny  []string `json:"deny"`
}

// Func to parse CIDR string to net.IPNet
func ParseCIDR(cidr string) *net.IPNet {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}

	return ipnet
}

// Func to add IPNet to allow list
func (c *HttpAccessContext) AddIPNet(ipnet *net.IPNet) {
	c.Allow = append(c.Allow, ipnet.String())
}

// Func to add IPNet to deny list
func (c *HttpAccessContext) AddDenyIPNet(ipnet *net.IPNet) {
	c.Deny = append(c.Deny, ipnet.String())
}

// Func to add IP to allow list
func (c *HttpAccessContext) AddIP(ip string) {
	c.Allow = append(c.Allow, ip)
}

// Func to add IP to deny list
func (c *HttpAccessContext) AddDenyIP(ip string) {
	c.Deny = append(c.Deny, ip)
}

// Func to add CIDR to allow list
func (c *HttpAccessContext) AddCIDR(cidr string) {
	c.Allow = append(c.Allow, cidr)
}

// Func to add CIDR to deny list
func (c *HttpAccessContext) AddDenyCIDR(cidr string) {
	c.Deny = append(c.Deny, cidr)
}
