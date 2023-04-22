package main

import (
	"os"
	"text/template"
)

func main() {
	// We can create an in-line template using the Parse() method of the
	// template.Template type
	tmpl, err := template.New("simpleTemplate").Parse(`{{ "foobar" | print }}`)
	if err != nil {
		panic(err)
	}

	// Create the file to contain our output
	file, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}

	// Anything that satisfies io.Writer can be passed as the first parameter
	// to Execute(), which includes the "file" returned to us by os.Create
	err = tmpl.Execute(file, nil)
	if err != nil {
		panic(err)
	}
}
