package main

import "fmt"

func main() {

	// Simple loop which prints VLAN IDs in order from 1-5. It
	// accomplishes this by incrementing a counter `vlanID`
	// and exits once it's greater than 5
	vlanID := 1
	for vlanID <= 5 {
		fmt.Printf("VLAN %d\n", vlanID)
		vlanID = vlanID + 1
	}

	// The canonical version which declares and increments counter
	// while also stating exit condition - all on one line.
	for vlanID := 1; vlanID <= 5; vlanID++ {
		fmt.Printf("VLAN %d\n", vlanID)
	}

	// We can use the "continue" keyword to cause the loop to repeat
	// earlier than it normally would
	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue
		}

		// Because of the "continue" statement above,
		// this line will not run when i == 3
		fmt.Printf("VLAN %d\n", i)
	}

	vlanID = 1
	// Loops with no exit condition will loop indefinitely
	// unless a `break` or `return` statement is used
	for {

		// This line will always execute as long as the loop
		// is running
		fmt.Printf("Looking at VLAN %d\n", vlanID)

		// Because our loop doesn't have an exit condition,
		// this break statement is the only way this loop
		// will end
		if vlanID > 5 {
			break
		}

		vlanID++
	}

	// NESTED LOOPS
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

	// This label "deviceloop" applies to the outer loop which is declared
	// immediately below it. `continue` or `break` statements at any level
	// of nested loop to refer explicitly to this outer loop scope.
deviceloop:
	// Iterate through a slice of devices
	for _, device := range devices {
		// Iterate through a slice of that device's interfaces
		for i, iface := range device.interfaces {
			// Iterate through a slice of that interface's vlan IDs
			for _, vlanID := range iface.vlans {

				// We've found VLAN 400 - time to print the device and interface name, and break out of
				// all loops
				if vlanID == 400 {
					fmt.Printf("Device %s has vlan 400 configured on interface %d\n", device.hostname, i)

					// A normal "break" statement would only break out of the
					// inner loop - but by referring to the "deviceloop" label
					// we declared earlier, we can specify that we want to
					// break out of the outmost loop.
					break deviceloop
				}
			}
		}
	}

}
