package main

import "fmt"

func main() {

	// This declares "arr" as an array of type "int"
	// and a size of 3.
	var arr [3]int

	// Once initialized, we can set values in the array
	// by their index. Since arrays have a fixed size
	// the compiler can warn us if we use an invalid index
	//
	// Don't forget, slices and arrays start with index 0!
	arr[0] = 1

	// You can also initialize arrays with values at
	// the same time
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Initializing a slice is very similar to initializing
	// an array - just leave out the size!
	//
	// Note though that this is slice is empty - we will need to
	// append values to it before we can do much with it.
	var intSlice []int
	fmt.Println(intSlice)

	// You can create slices of just about any type - here's a slice
	// of strings!
	var stringSlice []string
	fmt.Println(stringSlice)

	// The "literal" method using curly braces also allows us to initialize
	// the slice with some values at the same time.
	var intSlice2 = []int{11, 22, 33, 44, 55}
	fmt.Println(intSlice2)

	// This is identical to the above example.
	intSlice3 := []int{11, 22, 33, 44, 55}
	fmt.Println(intSlice3)

	ourSlice := []int{11, 22, 33, 44, 55}

	// output: ourSlice cap is 5, len is 5
	//
	// The "cap()" function returns an integer containing the slice's capacity,
	// len() returns the slice's length. We can see that after initialization,
	// both are set to 5, meaning that the backing array has a capacity of 5,
	// and the "segment" of that backing array which the slice provides a view
	// to is also 5.
	fmt.Printf("ourSlice cap is %d, len is %d\n", cap(ourSlice), len(ourSlice))

	// append() takes the original slice, and a new element, and returns a new slice.
	// That's why we're passing "ourSlice" as the first parameter, but then overwriting
	// it with the result.
	ourSlice = append(ourSlice, 66)

	// output: ourSlice cap is 10, len is 6
	//
	// After appending a value, the slice length increased to 6
	// as expected, but the capacity is now 10! This is because we reached
	// the maximum capacity of the backing array, so append() had to allocate
	// a new one.
	fmt.Printf("ourSlice cap is %d, len is %d\n", cap(ourSlice), len(ourSlice))

	// Append one more time
	ourSlice = append(ourSlice, 77)

	// output: ourSlice cap is 10, len is 7
	//
	// After another append, the length has yet again increased to 7, but the
	// capacity remains unchanged, because it is greater than the length.
	// This means that append() did not have to allocate a new backing array,
	// it had enough room to spare to accommodate the additional element.
	fmt.Printf("ourSlice cap is %d, len is %d\n", cap(ourSlice), len(ourSlice))

	////

	// We can get the flexibility benefits of slices and the predictability/performance
	// of arrays by using make() to declare slices with a length (and capacity) ahead of time.
	largeCapSlice := make([]int, 2, 50)

	// output: largeCapSlice cap is 50, len is 2
	fmt.Printf("largeCapSlice cap is %d, len is %d\n", cap(largeCapSlice), len(largeCapSlice))

	// Because our length is 2, we can set the first two elements by index:
	largeCapSlice[0] = 1
	largeCapSlice[1] = 2

	// Beyond this, we must use append() - but since our slice has a capacity of 50, we can add 48
	// more elements before append() must allocate a new backing array. Efficient!
	for i := 3; i <= 50; i++ {
		largeCapSlice = append(largeCapSlice, i)
	}

	// output: largeCapSlice cap is 50, len is 50
	fmt.Printf("largeCapSlice cap is %d, len is %d\n", cap(largeCapSlice), len(largeCapSlice))

	//

	var intSliceIter = []int{11, 22, 33, 44, 55}

	// You can use a "for" loop to iterate over the slice using a counter variable.
	// Starting at 0 and ending before you reach the end of the slice allows you to
	// iterate over each element one at a time.
	for i := 0; i < len(intSliceIter); i++ {
		fmt.Printf("intSliceIter index %d has a value of %d\n", i, intSliceIter[i])
	}

	// Alternatively, you can use the "range" keyword to do the same thing. At each
	// iteration, the variable `i` will be set to the next index of the slice.
	for i := range intSliceIter {
		fmt.Printf("intSliceIter index %d has a value of %d\n", i, intSliceIter[i])
	}

	// `range` can also provide you with the value at each index.
	for i, val := range intSliceIter {
		fmt.Printf("intSliceIter index %d has a value of %d\n", i, val)
	}

	// When searching an array for a particular value, you can use the `break` statement
	// to stop iterating once you've found it.
	toFind := 33
	for i, val := range intSliceIter {
		if val == toFind {
			fmt.Printf("Found! Index is %d\n", i)

			// Since we've found our value, there's no point in looping any further.
			// We can use `break` to stop iterating over the slice.
			break
		}
	}

}
