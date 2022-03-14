package main

import (
	"os"
	"text/template"
)

func main() {

	switches := map[string]int{
		"sw01": 48,
		"sw02": 24,
		"sw03": 48,
	}

	// Since we're now iterating over a map, the two created variables after the "range" keyword
	// represent each key/value pair
	tmplStr := `
{{range $hostname, $ifCount := .}}
Device {{$hostname}} has {{$ifCount}} interfaces
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
