package main

import (
	"os"
	"text/template"
)

func main() {
	// First, we must define the struct that represents a switch
	type Switch struct {
		Hostname       string
		InterfaceCount uint
	}

	// Then we instantiate this struct into a variable "sw01"
	sw01 := Switch{"sw01", 48}

	// We can refer to the fields of the struct using the ".<field name>" syntax.
	tmpl, err := template.New("switchTemplate").Parse("Device {{.Hostname}} has {{.InterfaceCount}} interfaces\n")
	if err != nil {
		panic(err)
	}

	// Our input struct is passed in as the second parameter to Execute()
	err = tmpl.Execute(os.Stdout, sw01)
	if err != nil {
		panic(err)
	}
}
