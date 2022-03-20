package main

import (
	"os"
	"text/template"
)

func main() {
	type Switch struct {
		Hostname       string
		InterfaceCount uint
	}

	// switches is a slice that represents all our Switch instances
	switches := []Switch{
		{"sw01", 48},
		{"sw02", 24},
		{"sw03", 48},
	}

	// As with Jinja, it's often better to define templates in a multi-line string like "tmplStr" below,
	// or in a separate file that is read in before rendering.
	tmplStr := `
{{range $i, $switch := .}}
Device {{$switch.Hostname}} has {{$switch.InterfaceCount}} interfaces
{{end}}
`
	tmpl, err := template.New("switchTemplate").Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, switches)
	if err != nil {
		panic(err)
	}
}
