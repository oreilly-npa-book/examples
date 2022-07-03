package main

import "fmt"

func main() {

	// This is a boolean variable that we're setting explicitly
	// here, but this could be set any number of ways, such as in
	// response to parsing a config file.
	var snmpConfigured bool = true

	// "if" statements in Go work by testing whether an expression
	// evaluates to a boolean true or false.
	//
	// Because `snmpConfigured` is itself a boolean type, it can be
	// used as an expression.
	if snmpConfigured {
		// Code within the braces, will execute if the expression is true
		fmt.Println("SNMP is configured!")
	}

	// In this case, the "!" negates the value of snmpConfigured,
	// so the inner statement will only execute if snmpConfigured
	// is false.
	if !snmpConfigured {
		fmt.Println("SNMP is not configured.")
	}

	// Both conditions can be handled by using the `else`
	// keyword.
	if snmpConfigured {
		fmt.Println("SNMP is configured!")
	} else {
		fmt.Println("SNMP is not configured.")
	}

	// Boolean values can only be true or false. Other types, like
	// integers, can have a wider variety of possibilities.
	var vlanID int = 1024

	// To handle this, we can incorporate more complex expressions
	// using greater-than (>) and less-than (<) operators
	// in conjunction with "else if" statements
	if vlanID < 100 {
		fmt.Println("VLAN ID is less than 100")
	} else if vlanID > 100 && vlanID < 1000 {
		fmt.Println("VLAN ID is between 100 and 1000")
	} else {
		fmt.Println("VLAN ID is greater than 1000")
	}
}
