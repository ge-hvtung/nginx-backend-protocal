package httpcore_test

import (
	"fmt"
	"testing"

	"github.com/tufanbarisyildirim/gonginx/parser"
	httpcore "github.com/tunghauvan/nginx-backend-protocal/packages/nginx/http_core"
)

// test ParseErrorPageDirective
func TestParseErrorPageDirective(t *testing.T) {
	p := parser.NewStringParser(`
http {
	error_page 404 /404.html;
	error_page 500 502 503 504 /50x.html;
}`)
	conf := p.Parse()

	error_pages := []httpcore.ErrorPageContext{}

	all_error_page_directives := conf.FindDirectives("error_page")
	for _, error_page_directive := range all_error_page_directives {
		error_page_context, err := httpcore.ParseErrorPageDirective(error_page_directive)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", error_page_context)
		error_pages = append(error_pages, error_page_context)
	}

	// Assert output vs expected
	expected_error_pages := []httpcore.ErrorPageContext{
		{
			Codes: []int{404},
			URI:   "/404.html",
		},
		{
			Codes: []int{500, 502, 503, 504},
			URI:   "/50x.html",
		},
	}

	for i, error_page := range error_pages {
		if error_page.Codes[0] != expected_error_pages[i].Codes[0] {
			panic("Error code not match")
		}
		if error_page.URI != expected_error_pages[i].URI {
			panic("URI not match")
		}
	}

}
