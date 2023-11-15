package core

import (
	"fmt"
	"strconv"

	gonginx "github.com/tufanbarisyildirim/gonginx"
)

// Parse error_page directive from IDirective
func ParseErrorPageDirective(directive gonginx.IDirective) (errorPageContext ErrorPageContext, err error) {
	// Parse error_page
	errorPageContext = ErrorPageContext{}
	errorPageContext.Codes = []int{}
	errorPageContext.Response = ""
	errorPageContext.URI = ""

	// Print error_page directive
	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

	// Parse this directive "error_page 500 502 503 504 /50x.html;"
	// Return error if not starts with error_page
	if directive.GetName() != "error_page" {
		err = fmt.Errorf("ErrorPageContext: directive is not error_page")
		return errorPageContext, err
	}

	// Get arguments
	arguments := directive.GetParameters()

	fmt.Printf("Arguments: %v\n", arguments)

	// Get codes
	for _, argument := range arguments {
		// Check if string is a number
		if isNumber(argument) {
			// Convert string to int
			code, _ := stringToInt(argument)
			errorPageContext.Codes = append(errorPageContext.Codes, code)
			continue
		}

		// Check if string starts with /
		if string(argument[0]) == "/" {
			errorPageContext.URI = argument
			continue
		}

		// Check if argument is a @fallback
		if string(argument[0]) == "@" {
			errorPageContext.URI = argument
			continue
		}

		// Check if argument starts with =
		if string(argument[0]) == "=" {
			errorPageContext.Response = argument
			continue
		}

		// return error if not starts with /, @, = or number
		err = fmt.Errorf("ErrorPageContext: argument is not a number, starts with /, @, =")
		return errorPageContext, err
	}

	return errorPageContext, nil
}

func ParseHttpDirective(directive gonginx.IDirective) (httpContext HttpContext, err error) {
	// Print http directive
	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

	// Parse this directive "http { ... }"
	// Return error if not starts with http
	if directive.GetName() != "http" {
		err = fmt.Errorf("HttpContext: directive is not http")
		return httpContext, err
	}

	// Create a new httpContext
	httpContext = HttpContext{}

	// Get directives
	directives := directive.GetBlock().FindDirectives("server")

	// Parse directives
	for _, directive := range directives {
		// Parse server directive
		if directive.GetName() == "server" {
			// Parse server directive
			serverContext, err := ParseServerDirective(directive)
			if err != nil {
				return httpContext, err
			}

			// Add serverContext to httpContext
			httpContext.Servers = append(httpContext.Servers, &serverContext)
		}
	}

	return httpContext, nil
}

func ParseServerDirective(directive gonginx.IDirective) (serverContext ServerContext, err error) {
	// Print server directive
	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

	// Parse this directive "server { ... }"
	// Return error if not starts with server
	if directive.GetName() != "server" {
		err = fmt.Errorf("ServerContext: directive is not server")
		return serverContext, err
	}

	// Create a new serverContext
	serverContext = ServerContext{}

	// Get server_names
	server_names := directive.GetBlock().FindDirectives("server_name")
	serverContext.ServerNames = []string{}
	for _, server_name := range server_names {
		serverContext.ServerNames = append(serverContext.ServerNames, server_name.GetParameters()[0])
	}

	// Get directives
	listen_directives := directive.GetBlock().FindDirectives("listen")
	location_directives := directive.GetBlock().FindDirectives("location")

	directives := append(listen_directives, location_directives...)

	// Parse directives
	for _, directive := range directives {
		// Parse listen directive
		if directive.GetName() == "listen" {
			// Parse listen directive
			listenContext, err := ParseListenDirective(directive)
			if err != nil {
				return serverContext, err
			}

			// Add listenContext to serverContext
			serverContext.Listens = append(serverContext.Listens, listenContext)
		}
	}

	return serverContext, nil
}

func ParseListenDirective(directive gonginx.IDirective) (listenContext ListenContext, err error) {
	// Print listen directive
	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

	// Parse this directive "listen 80;"
	// Return error if not starts with listen
	if directive.GetName() != "listen" {
		err = fmt.Errorf("ListenContext: directive is not listen")
		return listenContext, err
	}

	// Create a new listenContext
	listenContext = ListenContext{}

	// Get arguments
	arguments := directive.GetParameters()

	for _, argument := range arguments {
		// Check if string is a number
		if isNumber(argument) {
			listenContext.Ports = append(listenContext.Ports, argument)
			continue
		}

		// Check argument is ssl
		if argument == "ssl" {
			listenContext.SSL = true
			continue
		}

		// Other arguments map to listen
		listenContext.Listen = append(listenContext.Listen, argument)
		// return error if not starts with /, @, = or number
		// err = fmt.Errorf("ListenContext: argument is not a number, starts with /, @, =")
		return listenContext, err
	}

	return listenContext, nil
}

// func ParseLocationDirective(directive gonginx.IDirective) (locationContext LocationContext, err error) {
// 	// Print location directive
// 	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

// 	// Parse this directive "location / { ... }"
// 	// Return error if not starts with location
// 	if directive.GetName() != "location" {
// 		err = fmt.Errorf("LocationContext: directive is not location")
// 		return locationContext, err
// 	}

// 	// Create a new locationContext
// 	locationContext = LocationContext{}

// 	// Get arguments
// 	arguments := directive.GetParameters()

// 	// Get path
// 	locationContext.Path = arguments[0]

// 	// Get directives
// 	directives := directive.GetDirectives()

// 	// Parse directives
// 	for _, directive := range directives {
// 		// Parse proxy_pass directive
// 		if directive.GetName() == "proxy_pass" {
// 			// Parse proxy_pass directive
// 			proxyPassContext, err := ParseProxyPassDirective(directive)
// 			if err != nil {
// 				return locationContext, err
// 			}

// 			// Add proxyPassContext to locationContext
// 			locationContext.ProxyPass = proxyPassContext
// 		}

// 		// Parse error_page directive
// 		if directive.GetName() == "error_page" {
// 			// Parse error_page directive
// 			errorPageContext, err := ParseErrorPageDirective(directive)
// 			if err != nil {
// 				return locationContext, err
// 			}

// 			// Add errorPageContext to locationContext
// 			locationContext.ErrorPages = append(locationContext.ErrorPages, errorPageContext)
// 		}

// 		// Parse access directive
// 		if directive.GetName() == "access" {
// 			// Parse access directive
// 			accessContext, err := ParseAccessDirective(directive)
// 			if err != nil {
// 				return locationContext, err
// 			}

// 			//
// 			locationContext.Access = accessContext

// 		}

// 		// Parse return directive
// 		if directive.GetName() == "return" {

// 		}
// 	}

// 	return locationContext, nil
// }

// func ParseProxyPassDirective(directive gonginx.IDirective) (proxyPassContext ProxyPassContext, err error) {
// 	// Print proxy_pass directive
// 	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

// 	// Parse this directive "proxy_pass http://localhost:8080;"
// 	// Return error if not starts with proxy_pass
// 	if directive.GetName() != "proxy_pass" {
// 		err = fmt.Errorf("ProxyPassContext: directive is not proxy_pass")
// 		return proxyPassContext, err
// 	}

// 	// Create a new proxyPassContext
// 	proxyPassContext = ProxyPassContext{}

// 	// Get arguments
// 	arguments := directive.GetParameters()

// 	// Get url
// 	proxyPassContext.URL = arguments[0]

// 	return proxyPassContext, nil
// }

// func ParseAccessDirective(directive gonginx.IDirective) (accessContext AccessContext, err error) {
// 	// Print access directive
// 	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

// 	// Parse this directive "access { allow
// 	//
// 	// Return error if not starts with access

// 	// Create a new accessContext
// 	accessContext = AccessContext{}

// 	// Get directives
// 	directives := directive.GetDirectives()

// 	// Parse directives
// 	for _, directive := range directives {
// 		// Parse allow directive
// 		if directive.GetName() == "allow" {
// 			// Parse allow directive
// 			allowContext, err := ParseAllowDirective(directive)
// 			if err != nil {
// 				return accessContext, err
// 			}

// 			// Add allowContext to accessContext
// 			accessContext.Allow = allowContext
// 		}

// 		// Parse deny directive
// 		if directive.GetName() == "deny" {
// 			// Parse deny directive
// 			denyContext, err := ParseDenyDirective(directive)
// 			if err != nil {
// 				return accessContext, err
// 			}

// 			// Add denyContext to accessContext
// 			accessContext.Deny = denyContext
// 		}
// 	}

// 	return accessContext, nil
// }

// func ParseAllowDirective(directive gonginx.IDirective) (allowContext AllowContext, err error) {
// 	// Print allow directive
// 	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

// 	// Parse this directive "allow
// 	//
// 	// Return error if not starts with allow

// 	// Create a new allowContext
// 	allowContext = AllowContext{}

// 	// Get arguments
// 	arguments := directive.GetParameters()

// 	// Get ip
// 	allowContext.IP = arguments[0]

// 	return allowContext, nil
// }

// func ParseDenyDirective(directive gonginx.IDirective) (denyContext DenyContext, err error) {
// 	// Print deny directive
// 	fmt.Println(gonginx.DumpDirective(directive, gonginx.IndentedStyle))

// 	// Parse this directive "deny
// 	//
// 	// Return error if not starts with deny

// 	// Create a new denyContext
// 	denyContext = DenyContext{}

// 	// Get arguments
// 	arguments := directive.GetParameters()

// 	// Get ip
// 	denyContext.IP = arguments[0]

// 	return denyContext, nil
// }

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
