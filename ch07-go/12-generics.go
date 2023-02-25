package main

import (
	"fmt"
)

func main() {

	// Because Min uses generic parameters, we can pass a variety of types
	// as parameters to x and y - in this example we're passing integers, then
	// floats, then strings. Because all of these are listed in the type set declared
	// in the "comparable" interface, this works
	fmt.Println(Min(3, 5))
	fmt.Println(Min(2.5, 6.3))
	fmt.Println(Min("foo", "fooooo"))
}

// Rather than declaring methods in our "comparable" interface, we can specify some types
// which we know can be compared using the < operator. This allows us to pass any
// of these types into Min(). This is the upgrade that interfaces received in
// Go 1.18 - they can be used to define method sets and/or type sets.
type comparable interface {
	int | float64 | string
}

// Here, we're declaring a generic type T which must implement the "comparable" interface.
// Then, when declaring parameters x and y, we can reference this type T. This means that both
// x and y must in turn implement "comparable". Also, because they're both generic type T, they
// cannot be different types from each other when the function is invoked.
//
// Without generics, we'd either have to use interfaces (and therefore each of these types
// would have to have methods the interface would match on) or we'd have to create a copy
// of this Min function for each type we wanted to be able to pass in (Min Int, MinFloat64, etc)
func Min[T comparable](x, y T) T {
	if x < y {
		return x
	}
	return y
}
