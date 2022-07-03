package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	doPrint()

	printMessage("Go")
	// output: Hello network automators, today we're learning Go!

	// Create a variable to hold the prefix length we'll pass in to totalIPv4Addresses()
	prefixLen := 22
	// Call totalIPv4Addresses and provide the variable prefixLen as the required parameter.
	// (we could pro ide a value of 22 directly as well ,but this way we can re-use this variable
	// in the log message below).
	//
	// Note also that we're assigning the return value to a new variable called "addrs"
	addrs := totalIPv4Addresses(prefixLen)
	fmt.Printf("There are %d addresses in a /%d prefix\n", addrs, prefixLen)
	// output: There are 1024 addresses in a /22 prefix

	// both input parameters and return values are separated by commas
	sum, product := sumAndProduct(2, 6)
	fmt.Printf("Sum is %d, product is %d\n", sum, product)

	// It is conventional to assign error return types to a variable "err".
	err := createVLAN(50)
	// If err is not a `nil` value, it means an error occurred, so we should
	// check for that immediately after calling the function above.
	if err != nil {
		// This is where you could take steps to recover from the error
		// if possible.
		fmt.Println(err)
	}

}

func doPrint() {
	fmt.Println("Hello network automators!")
	fmt.Println("Welcome to Network Programmability and Automation!")
	fmt.Println("Enjoy this chapter on the Go programming language.")
}

func printMessage(msg string) {
	fmt.Printf("Hello network automators, today we're learning %s!\n", msg)
}

// totalIPv4Addresses is a function for calculating the number of addresses
// in an IPv6 prefix of a provided length.
//
// `prefixLen` is a parameter of type `int`.
// The return type for this function is also `int`, which is declared after the
// set of parentheses which contain the function parameter(s).
func totalIPv4Addresses(prefixLen int) int {

	// To calculate the number of addresses in an IPv6 address, we must
	// calculate 2^x, where x is prefixLen subtracted from 128. So
	// let's first get x.
	x := 32 - prefixLen

	// Go doesn't have an exponent operator, so we will use the `Pow()`
	// function in the `math` package. This function has two parameters,
	// and each have a type of `float64`, which is why we're converting
	// the latter using the `float64` built-in (2.0 is already a float64).
	addrCount := math.Pow(2.0, float64(x))

	// `math.Pow()` also has a return type of `float64`, so we must convert it to
	// an `int` before returning it, to satisfy our function's return type.
	// The "return" keyword allows us to exit a function immediately, and return
	// the value provided.
	return int(addrCount)
}

// sumAndProduct takes in two integers x and y, and returns their sum and product, respectively.
//
// Note that the input parameters x and y are separated only by a comma - since they're both integers,
// we can just specify `int` once after them.
//
// Also, note that the return types are also separated by a comma and also wrapped in parentheses.
func sumAndProduct(x, y int) (int, int) {
	sum := x + y
	product := x * y
	return sum, product
}

// createVLAN takes in an unsigned integer parameter for the VLAN ID, and
// returns an error type if a problem is encountered.
func createVLAN(id uint) error {

	// Even though the `uint` type (used for the `id` parameter) can support billions
	// of values, we know that 4096 is the maximum VLAN ID. So we can add
	// a conditional which checks for this, and returns a new error if the ID is
	// over this value.
	if id > 4096 {

		// New() is a function in the `errors` package which allows us to initialize a
		// new error value from a string containing our custom error message.
		return errors.New("VLAN ID must be <= 4096")
	}

	// This will only execute if `id` is a valid VLAN ID, so to simulate the creation
	// of a VLAN, we'll print a log message and return `nil` as our error value, which
	// indicates no error occurred.
	fmt.Printf("Creating VLAN with ID of %d\n", id)
	return nil
}

func returnFromNestedLoops() {
	type Interface struct {
		vlans []int
	}

	type Device struct {
		hostname   string
		interfaces []Interface
	}

	devices := []Device{
		{
			hostname: "sw01",
			interfaces: []Interface{
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
			},
		},
		{
			hostname: "sw02",
			interfaces: []Interface{
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300, 400},
				},
				{
					vlans: []int{100, 200, 300},
				},
			},
		},
		{
			hostname: "sw03",
			interfaces: []Interface{
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
				{
					vlans: []int{100, 200, 300},
				},
			},
		},
	}

	// Iterate through a slice of devices
	//
	// Note that no label is needed here, as we're using the `return` keyword
	// instead!
	for _, device := range devices {
		// Iterate through a slice of that device's interfaces
		for i, iface := range device.interfaces {
			// Iterate through a slice of that interface's vlan IDs
			for _, vlanID := range iface.vlans {
				// We've found VLAN 400 - time to print the device and interface name, and break out of
				// all loops
				if vlanID == 400 {
					fmt.Printf("Device %s has vlan 400 configured on interface %d\n", device.hostname, i)
					// we can break out of any nested level by simply returning from
					// the containing function.
					return
				}
			}
		}
	}
}
