package main

import "fmt"

func main() {

	// Simple loop which increments a counter
	// and exits once it's greater than 5
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// The canonical version which declares and increments counter
	// while also stating exit condition - all on one line.
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// We can use the "continue" keyword to cause the loop to repeat
	// earlier than it normally would
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		// Because of the "continue" statement above,
		// this line will not run when i == 3
		fmt.Println(i)
	}

	j := 1
	// Loops with no exit condition will loop indefinitely
	// unless a `break` or `return` statement is used
	for {

		// This line will always execute as long as the loop
		// is running
		fmt.Println("The value of j is", j)

		// Because our loop doesn't have an exit condition,
		// this break statement is the only way this loop
		// will end
		if j > 5 {
			break
		}
		j++
	}

	// counter tracks the number of times the inner loop has executed.
	// It's declared outside either loop so that it will retain its value
	counter := 1

	// This label "outerloop" applies to the outer loop which is declared
	// immediately below it. `continue` or `break` statements can use this
	// label to refer to this loop scope.
outerloop:

	// This outer loop has no exit condition, so it will
	// loop infinitely unless we break out of it
	for {

		// This inner loop has an exit condition, but once it finishes,
		// the outer loop will repeat, causing this inner loop to also
		// repeat
		for i := 1; i <= 3; i++ {

			if counter > 9 {

				// A normal "break" statement would only break out of the
				// inner loop - but by referring to the "outerloop" label
				// we declared earlier, we can specify that we want to
				// break out of the outer loop.
				break outerloop
			}
			fmt.Println("counter has a value of", counter)
			counter++
		}
	}

}
