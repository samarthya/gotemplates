package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/samarthya/learn-templates/four"
	"github.com/samarthya/learn-templates/one"
	"github.com/samarthya/learn-templates/three"
	"github.com/samarthya/learn-templates/two"
)

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
}
