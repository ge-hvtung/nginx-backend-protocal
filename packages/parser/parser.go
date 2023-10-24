package parser

import (
	"fmt"
	"strings"

	"github.com/tunghauvan/nginx-backend-protocal/internal/models"

	"github.com/tufanbarisyildirim/gonginx"

	nginxparser "github.com/tufanbarisyildirim/gonginx/parser"
)

type NgxParser interface {
	InitParse() *gonginx.Config
	SetConfig(config string)

	ParseUpstreams() ([]models.NgxUpstream, error)
	ParseServers() ([]models.NgxServer, error)
}

func NewNgxParser() NgxParser {
	return &NgxParserImpl{}
}

type NgxParserImpl struct {
	config       string
	nginx_config *gonginx.Config
}

func (p *NgxParserImpl) SetConfig(config string) {
	p.config = config
}

func (p *NgxParserImpl) InitParse() *gonginx.Config {
	pser := nginxparser.NewStringParser(p.config)
	nginx_config := pser.Parse()
	p.nginx_config = nginx_config
	return nginx_config
}

// Func to parse all upstreams
func (p *NgxParserImpl) ParseUpstreams() ([]models.NgxUpstream, error) {
	// Init parser
	conf := p.InitParse()
	// fmt.Println(gonginx.DumpBlock(conf.Block, gonginx.IndentedStyle))

	// Get all upstreams
	upstreams := conf.FindUpstreams()
	// Print upstreams
	fmt.Println(upstreams)

	// Create a slice of upstreams
	ngxUpstreams := make([]models.NgxUpstream, 0)

	// Loop through upstreams
	for _, upstream := range upstreams {
		// Create a new upstream
		ngxUpstream := models.NgxUpstream{
			UpstreamName: upstream.UpstreamName,
		}

		// Get all servers
		upstreamServer := upstream.UpstreamServers

		// Create a slice of servers
		ngxUpstreamServers := make([]models.NgxUpstreamServer, 0)

		// Loop through servers
		for _, server := range upstreamServer {
			// Create a new server
			ngxUpstreamServer := models.NgxUpstreamServer{
				Address: server.Address,
			}

			// Append server to slice of servers
			ngxUpstreamServers = append(ngxUpstreamServers, ngxUpstreamServer)
		}

		// Append servers to upstream
		ngxUpstream.UpstreamServers = ngxUpstreamServers

		// Append upstream to slice of upstreams
		ngxUpstreams = append(ngxUpstreams, ngxUpstream)
	}

	return ngxUpstreams, nil
}

func (p *NgxParserImpl) GetDirectives(parent gonginx.IDirective) map[string][]string {
	directives := make(map[string][]string)
	// Loop through server directives
	for _, directive := range parent.GetBlock().GetDirectives() {
		// Add directive to map
		directives[directive.GetName()] = directive.GetParameters()
	}
	return directives
}

// Func to parse all servers
func (p *NgxParserImpl) ParseServers() ([]models.NgxServer, error) {
	// Init parser
	conf := p.InitParse()
	// fmt.Println(gonginx.DumpBlock(conf.Block, gonginx.IndentedStyle))

	// Get all servers
	servers := conf.Block.FindDirectives("server")

	// Create a slice of servers
	ngxServers := make([]models.NgxServer, 0)

	// Loop through servers
	for _, server := range servers {
		// Create map of server directives
		serverDirectives := p.GetDirectives(server)

		// Get server name from directive
		// Create a new server
		ngxServer := models.NgxServer{
			ServerName: serverDirectives["server_name"],
		}

		// Add proxy props to server
		// Get all proxy directives
		serverProxyDirectives, err := p.GetProxyProps(serverDirectives)
		if err != nil {
			return nil, err
		}
		ngxServer.ProxyProps = serverProxyDirectives

		// Get all locations
		locations := server.GetBlock().FindDirectives("location")

		// Create a slice of locations
		ngxLocations := make([]models.NgxLocation, 0)

		// Loop through locations
		for _, location := range locations {
			// Create map of location directives
			locationDirectives := p.GetDirectives(location)

			// Get all proxy directives
			locationProxyDirectives, err := p.GetProxyProps(locationDirectives)
			if err != nil {
				return nil, err
			}

			// Create a new location
			ngxLocation := models.NgxLocation{
				LocationPath: location.GetParameters()[0],
			}

			// Get all proxy_pass
			proxy_pass := location.GetBlock().FindDirectives("proxy_pass")

			// Loop through proxy_pass
			for _, proxy := range proxy_pass {
				// Set proxy_pass
				ngxLocation.LocationProxyPass = proxy.GetParameters()[0]
			}

			// Add proxy props to location
			ngxLocation.ProxyProps = locationProxyDirectives

			// Append location to slice of locations
			ngxLocations = append(ngxLocations, ngxLocation)

			// Append locations to server
			ngxServer.Locations = ngxLocations
		}

		// Get all includes
		includes := server.GetBlock().FindDirectives("include")

		// Create a slice of includes
		ngxIncludes := make([]string, 0)

		// Loop through includes
		for _, include := range includes {
			// Append include to slice of includes
			ngxIncludes = append(ngxIncludes, include.GetParameters()[0])
		}

		// Append includes to server
		ngxServer.Includes = ngxIncludes

		// Append server to slice of servers
		ngxServers = append(ngxServers, ngxServer)
	}
	return ngxServers, nil
}

// Function to get all directives starting with proxy_* from a server
func (p *NgxParserImpl) GetProxyProps(directives map[string][]string) ([]string, error) {
	fmt.Printf("Directives: %v\n", directives)
	// Create a slice of proxy directives
	proxyDirectives := make([]string, 0)

	// Loop through directives
	for directive, value := range directives {
		// Continue if directive is less than 6 characters
		if len(directive) < 6 {
			continue
		}

		// Skip if start with proxy_pass or proxy_redirect
		if len(directive) > 10 && directive[:10] == "proxy_pass" {
			continue
		}

		// Trim prefix from directive
		fmt.Printf("Directive: %v\n", directive)
		fmt.Printf("Value: %v\n", value)

		// Check if directive starts with proxy_
		if directive[:6] == "proxy_" {
			// Create full directive will all values
			full_directive := directive + " " + strings.Join(value, " ")
			proxyDirectives = append(proxyDirectives, full_directive)
		}
	}
	fmt.Printf("Proxy Directives: %v\n", proxyDirectives)
	return proxyDirectives, nil

}

// Function to split a header into a slice of strings
func splitHeader(header string) []string {
	// Split header
	headerSplit := make([]string, 2)
	headerSplit[0] = header[:len(header)-len(header[strings.Index(header, " "):])]
	headerSplit[1] = header[strings.Index(header, " ")+1:]
	return headerSplit
}

// Function for parsing specific location directives
func (p *NgxParserImpl) ParseLocationDirectives(location *gonginx.Directive) (models.NgxLocation, error) {
	// Create a new location
	ngxLocation := models.NgxLocation{
		LocationPath: location.GetParameters()[0],
	}

	// Create map of location directives
	locationDirectives := make(map[string][]string)

	// Loop through location directives
	for _, directive := range location.GetBlock().GetDirectives() {
		// Add directive to map
		locationDirectives[directive.GetName()] = directive.GetParameters()
	}

	// Get all proxy directives
	locationProxyDirectives, err := p.GetProxyProps(locationDirectives)
	if err != nil {
		return models.NgxLocation{}, err
	}

	// Add proxy props to location
	ngxLocation.ProxyProps = locationProxyDirectives

	return ngxLocation, nil
}

// Function for parsing error_page location directives
// error_page 404 /404.html;
// location = /40x.html {
// }
