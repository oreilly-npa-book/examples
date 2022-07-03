package main

import "fmt"

func main() {

	// This is the most explicit form - it initalizes the
	// variable with a value, while also explicitly stating
	// its type.
	var myString1 string = "Hello, network gophers!"

	// Omitting the value defaults to the "zero value" for
	// that type - in this case, a zero-length string ("").
	// Other types have different zero values; for instance,
	// a boolean will default to false, and an int type will
	// default to 0
	var myString2 string

	// The Go compiler will infer the type based on the value
	// passed in.
	var myString3 = "Hello, network gophers!"

	// The := operator is short-hand for the previous example,
	// which infers the type based on the provided value.
	myString4 := "Hello!"

	// The var keyword allows you to group variable declarations
	// together, to improve readability
	var (
		myString5 string
		myString6 = "Hello, network gophers!"
	)

	// This is invalid, as there is no explicit type declaration
	// or a value from which the type can be inferred.
	// So, this will fail to compile.
	// var whatIsThis

	// Our program will fail to compile unless we use the variables
	// we've declared. Passing them to fmt.Println is an easy way
	// to get around this, but any usage should suffice
	fmt.Println(myString1)
	fmt.Println(myString2)
	fmt.Println(myString3)
	fmt.Println(myString4)
	fmt.Println(myString5)
	fmt.Println(myString6)

}
