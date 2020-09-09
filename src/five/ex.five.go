package five

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/samarthya/learn-templates/util"
)

const (
	// Filename It will be used for debugging
	Filename = "five"
)

// TemplateFile defines the contents of a template to be stored in a file, for testing.
type TemplateFile struct {
	// Identifies the name of the file
	Name string
	// Content of the file
	Contents string
}

// CreateTestDir Creates a temp directory.
func CreateTestDir(files []TemplateFile) string {
	util.Debug(Filename, "function called CreateTestDir")

	// TempDir creates a new temporary directory in the directory dir.
	dir, err := ioutil.TempDir("", "template")

	if err != nil {
		log.Fatal(err)
	}

	util.Debug(Filename, "Directory - "+dir+" created.")

	//No error
	for _, file := range files {
		f, err := os.Create(filepath.Join(dir, file.Name))
		if err != nil {
			log.Fatal(err)
		}
		util.Debug(Filename, "File opened "+file.Name)
		defer f.Close()
		_, err = io.WriteString(f, file.Contents)
		if err != nil {
			util.Fatal(err)
		}
	}
	return dir
}
