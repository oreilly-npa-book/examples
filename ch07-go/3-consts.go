package main

func main() {

	// Most explicit form, which specifies the type and the "const" keyword
	const myString string = "Hello, network gophers!"

	// This will fail to compile - myString is declared above as a const,
	// so a re-assignment like this is not permitted
	// myString = "new value"

	// As with variables, the compiler will infer the type from the assigned value
	// if not explicitly stated.
	const Pi = 3.14159265358979323846264338327950288419716939937510582097494459

	// Multiple assignment blocks work the same as variables as well
	const (
		myString2 = "Hello, network gophers!"
		retries   = 3
	)
}
