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

	// We can render the template with "Execute", passing in "os.Stdout"
	// as the first parameter, so we can see the results in our terminal
	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
