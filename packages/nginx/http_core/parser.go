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

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func stringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
