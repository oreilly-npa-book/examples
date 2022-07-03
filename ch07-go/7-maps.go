package main

import "fmt"

func main() {

	// CAREFUL - this syntax will only declare the map, but will not
	// initialize it. Trying to write to this map will cause a runtime panic.
	// var myMap map[string]int
	// myMap["foo"] = 80

	// It's much safer to declare and initialize the map at the same time.
	// Each of these examples is equivalent - they each declare and initialize
	// a map with a "string" type for the keys, and an "int" type for the values.
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	var myMap2 = map[string]int{}
	fmt.Println(myMap2)

	myMap3 := map[string]int{}
	fmt.Println(myMap3)

	// The "literal" method using curly braces also allows us to initialize
	// the map with some values at the same time.
	vlanMap := map[string]int{
		"VLAN_100": 100,
		"VLAN_200": 200,
		"VLAN_300": 300,
	}
	fmt.Println(vlanMap)

	//

	// This reads a value from the map using the expected key
	// and creates a new variable "vlan" with this value
	vlan := vlanMap["VLAN_300"]
	fmt.Printf("vlan is %d\n", vlan)

	// This syntax adds a single key-value pair to the map.
	// Note that our key is a string, and the value is an int, which
	// matches the types declared when the map was created
	vlanMap["VLAN_400"] = 401

	// We can overwrite an existing key
	vlanMap["VLAN_400"] = 400

	// You can delete a key/value pair from a map using the delete() function
	// This requires two parameters - first the map itself, and then the key to delete.
	delete(vlanMap, "VLAN_300")
	fmt.Println(vlanMap)

	// Reading a key that doesn't exist will return the zero-value
	// for the value's type - in this case, 0.
	fmt.Println(vlanMap["VLAN_999"]) // output: 0

	//

	// The same syntax used for reading a key from a map can optionally return a second
	// boolean value, which is set to `found` below. We can then test if `found` is true
	// (which indicates the key already exists) on the same line by adding a semicolon
	// and then the variable `found` on its own (remember, booleans can be used as entire
	// expressions in conditionals)
	if val, found := vlanMap["VLAN_999"]; found {
		// The key doesn't exist, so this will execute.
		fmt.Printf("Found vlan %d\n", val)
	}

	// We can test the reverse by simply negating `found` - this is an easy
	// way to test if a key is **not** found in a map.
	//
	// In this case, we don't expect the retrieved value to be useful, so we
	// can ignore it by replacing `val` with an underscore. This tells the compiler
	// to discard the retrieved value.
	if _, found := vlanMap["VLAN_999"]; !found {
		// The key doesn't exist, so this WILL execute.
		fmt.Println("Did not find VLAN_999 in the map")
	}

	if val, found := vlanMap["VLAN_400"]; found {
		// This key does exist, so we will see this print statement execute.
		fmt.Printf("Found vlan %d\n", val)
	}

	//

	// Like we saw with slices, the range keyword allows us to easily
	// iterate over the key/value pairs in the map.
	// Note that unlike slices, maps are not ordered.
	for key, value := range vlanMap {
		fmt.Printf("%s has a value of %d\n", key, value)
	}

	// If we don't need the values, we can leave that out to only retrieve the
	// keys out of the map.
	for key := range vlanMap {
		fmt.Printf("Found key %s\n", key)
	}
}
