package main

import (
	"os"
	"text/template"
)

func main() {
	type Switch struct {
		Hostname       string
		InterfaceCount uint
		Enabled        bool
	}

	// switches is a slice that represents all our Switch instances
	switches := []Switch{
		{"sw01", 48, true},
		{"sw02", 24, true},
		{"sw03", 48, false},
	}

	tmplStr := `
{{range $i, $switch := .}}
{{if $switch.Enabled}}
Device {{$switch.Hostname}} has {{$switch.InterfaceCount}} interfaces
{{end}}
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
