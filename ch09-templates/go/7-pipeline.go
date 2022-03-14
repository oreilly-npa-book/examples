package main

import (
	"os"
	"text/template"
)

// Note that Interfaces is now a slice of strings
type Switch struct {
	Hostname   string
	Interfaces []string
}

func main() {

	// A fairly small switch, indeed
	sw01 := Switch{"sw01", []string{
		"ge-0/0/1",
		"ge-0/0/2",
		"ge-0/0/3",
		"ge-0/0/4",
	}}

	tmplStr := "Device {{.Hostname}} has {{ .Interfaces | len }} interfaces\n"

	tmpl, err := template.New("switchTemplate").Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sw01)
	if err != nil {
		panic(err)
	}
}
