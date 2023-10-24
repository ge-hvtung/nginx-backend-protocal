package httpcore

import (
	"reflect"

	"github.com/tufanbarisyildirim/gonginx"
	httpaccess "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_access"
)

type Directives struct {
	Directives []gonginx.IDirective
}

// Func add directive to Directives
func (d *Directives) AddDirective(directive gonginx.IDirective) {
	d.Directives = append(d.Directives, directive)
}

// Func add http_access  to Directives
func (d *Directives) AddHttpAccessContext(c httpaccess.HttpAccessContext) {
	// Add allow directive
	if len(c.Allow) > 0 {
		allow_directive := gonginx.Directive{
			Name:       "allow",
			Parameters: c.Allow,
		}

		d.Directives = append(d.Directives, &allow_directive)
	}

	// Add deny directive
	if len(c.Deny) > 0 {
		deny_directive := gonginx.Directive{
			Name:       "deny",
			Parameters: c.Deny,
		}

		d.Directives = append(d.Directives, &deny_directive)
	}
}

// Func add CoreProps to Directives from LocationContext o
func (d *Directives) AddCoreProps(v reflect.Value) {
	typeOfS := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface().(string) == "" {
			continue
		}

		// Create a directive from the coreProp with Name lowercased and split by _
		directive := gonginx.Directive{
			Name:       toLowerSnakeCase(typeOfS.Field(i).Name),
			Parameters: []string{v.Field(i).Interface().(string)},
		}

		// Append the directive to the directives
		d.Directives = append(d.Directives, &directive)
	}
}

func (d *Directives) AddErrorPageContext(c ErrorPageContext) {
	// Add error_page directive
	if c.URI != "" {
		params := []string{}
		params = append(params, intSliceToString(c.Codes))

		if c.Response != "" {
			params = append(params, c.Response)
		}

		params = append(params, c.URI)

		error_page_directive := gonginx.Directive{
			Name:       "error_page",
			Parameters: params,
		}

		d.Directives = append(d.Directives, &error_page_directive)
	}
}
