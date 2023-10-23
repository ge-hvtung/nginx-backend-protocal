package httpaccess

import "net"

type AllowContext struct {
	IPNets []*net.IPNet `json:"ip_nets"`
}

type DenyContext struct {
	IPNets []*net.IPNet `json:"ip_nets"`
}

type HttpAccessContext struct {
	Allow []*AllowContext `nginx:"allow"`
	Deny  []*DenyContext  `nginx:"deny"`
}

// func init HTTPAccessModule and check about context valid
type HttpAccessModule interface {
	// check parent context is valid type of httpContext or not
	// if valid, return httpContext
	// else return error

	

