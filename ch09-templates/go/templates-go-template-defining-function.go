package main

import (
	"os"
	"strconv"
	"strings"
	"text/template"
)

type Switch struct {
	Hostname   string
	Interfaces []string
}

type Interface struct {
	Speed int
	FPC   int
	PIC   int
	Port  int
}

func IfParse(ifaceStr string) Interface {

	iface := Interface{}

	ifSplit := strings.Split(ifaceStr, "-")

	speeds := map[string]int{
		"ge": 1,
		"xe": 10,
		"et": 40,
	}
	iface.Speed = speeds[ifSplit[0]]

	locSplit := strings.Split(ifSplit[1], "/")

	fpc, _ := strconv.Atoi(locSplit[0])
	iface.FPC = fpc

	pic, _ := strconv.Atoi(locSplit[1])
	iface.PIC = pic

	port, _ := strconv.Atoi(locSplit[2])
	iface.Port = port

	return iface
}

func main() {

	// Create a mapping of functions to names we can refer to in our template.
	fmap := template.FuncMap{"ifparse": IfParse}

	sw01 := Switch{"sw01", []string{
		"ge-0/0/1",
		"ge-0/0/2",
		"ge-0/0/3",
		"ge-0/0/4",
	}}

	// We're first creating a new variable "loc" which is created by
	// pipelining "$interface" into our custom function "ifparse".
	//
	// Since "ifparse" returns an "Interface" type, we can then refer to the fields
	// of that struct directly!
	tmplStr := `
{{range $i, $interface := .Interfaces}}
{{with $loc := $interface | ifparse}}
Interface {{$interface}}   port {{$loc.Port}}
{{end}}
{{end}}
`

	// Don't forget to pass in the FuncMap using the "Funcs" method, as shown here.
	tmpl, err := template.New("switchTemplate").Funcs(fmap).Parse(tmplStr)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sw01)
	if err != nil {
		panic(err)
	}
}
