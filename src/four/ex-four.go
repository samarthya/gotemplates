package four

import (
	"html/template"
	"log"
	"os"
	"strings"
)

const templateText = `
Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
Output 3: {{println .}}
`

// First we create a FuncMap with which to register the function.
var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text.
	// Look for the `templateText`
	"title": strings.Title,
}

// Example Method exposed
func Example(s string) {

	// A simple template definition to test our function.
	// We print the input text several ways:
	// - the original
	// - title-cased
	// - title-cased and then printed with %q
	// - printed with %q and then title-cased.

	// Create a template, add the function map, and parse the text.
	tmpl, err := template.New("titleTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	// Run the template to verify the output.
	err = tmpl.Execute(os.Stdout, s)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}
