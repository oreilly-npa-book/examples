// Packages allow you to organize your code into a logical hierarchy.
// We are specifying the "main" package here becuase we want to create
// an executable program. We'll explore packages in greater detail
// in later examples.
package main

// The "import" keyword allows us to use other packages in our code.
// In this case, the "fmt" package is part of the standard library
// and is used for formatted I/O.
import "fmt"

// The "main" function is the primary entrypoint when creating an executable
// program. This means that when we run a compiled Go program, this code
// represents the beginning of that program's logical flow.
func main() {

	// The `Println` function is part of the `fmt` package we imported above,
	// and allows us to print a simple string to the terminal as a line
	// of output.
	fmt.Println("Hello, network automators!")
}
