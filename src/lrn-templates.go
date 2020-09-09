package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/samarthya/learn-templates/five"
	"github.com/samarthya/learn-templates/four"
	"github.com/samarthya/learn-templates/one"
	"github.com/samarthya/learn-templates/three"
	"github.com/samarthya/learn-templates/two"
)

var dir string

func init() {
	// Here we create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir = five.CreateTestDir([]five.TemplateFile{
		// T0.tmpl is a plain template file that just invokes T1.
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}})`},
		// T1.tmpl defines a template, T1 that invokes T2.
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})

	log.Println(" Init called.")
}

func exampleFive() {

	fmt.Println("\n Example Five - \n")
	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// Here starts the example proper.
	// T0.tmpl is the first name matched, so it becomes the starting template,
	// the value returned by ParseGlob.
	tmpl := template.Must(template.ParseGlob(pattern))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}
}

// exampleThree
func exampleThree() {

	fmt.Println("\n Example Three - \n")

	var (
		// FuncMap is the type of the map defining the mapping from names to functions.
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)

	masterTmpl, err := template.New("master").Funcs(funcs).Parse(three.Master)

	// If parsing fails
	if err != nil {
		log.Fatal(err)
	}

	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(three.Overlay)

	if err != nil {
		log.Fatal(err)
	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}

// exampleTwo - Uses the actions and pipelines
func exampleTwo() {
	fmt.Println("\n Example Two - \n")
	var recipients = []two.Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(two.Letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		fmt.Println("\n *************************** \n")
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
	fmt.Println("\n ***************************\n ")

}

// exampleOne Example One, prints `17 items are made of wool.`
func exampleOne() {
	fmt.Println(" Example One - \n")
	sweaters := one.Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}.")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

// Main function
func main() {
	fmt.Printf(" *** Templates Learning *** \n")
	exampleOne()
	exampleTwo()
	exampleThree()
	four.Example("the go programming language")
	exampleFive()
}
