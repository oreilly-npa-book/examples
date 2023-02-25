package main

import (
	"errors"
	"fmt"
)

// This is where we define our custom struct `vlan`
// Note that this is just a definition - we'll actually create
// an instance of this later.
type vlan struct {

	// `id` and `name` are fields of our `vlan` struct.
	// They each have their own type definition (in this
	// case, uint and string)
	id   uint
	name string
}

// NewVLAN is a "constructor" function, which returns an instantiated
// `vlan` instance but also ensures that the `id` field is populated
// with a valid VLAN ID.
//
// Note that the first letter in this function is capitalized, indicating
// it is "exported" (accessible from outside the current package).
func NewVLAN(id uint, name string) (vlan, error) {
	if id > 4096 {

		// When returning a non-nil error alongside a struct type, it's
		// conventional to return the zero-value for those struct types.
		// You can do this with the empty braces as shown below.
		return vlan{}, errors.New("VLAN ID must be <= 4096")
	}

	// We've already determined that the `id` parameter satisfies our
	// requirements, so we can instantiate and return a `vlan` right
	// here in the return statement, passing in the `id` and `name`
	// variables as fields.
	return vlan{
		id,
		name,
	}, nil
}

type device struct {
	hostname string

	// Here, `vlans` is a field on the `device` struct. Its type
	// is a slice of `vlan` instances.
	vlans []vlan
}

// printHostname has no explicit parameters, but does have a receiver of type `device`
// named `d`.
func (d device) printHostname() {
	// We can refer to `d` in the body of the method
	// to access the fields of the instantiated struct object.
	fmt.Println(d.hostname)
}

// This has a value receiver so this won't work. Need a pointer receiver.
func (d device) setHostnameValueReceiver(hostname string) {
	d.hostname = hostname
}

// This has a **pointer receiver**, denoted by the asterisk before the `device` receiver type.
// This means that setting the `hostname` field here will apply to the original copy of this struct
// instance.
func (d *device) setHostname(hostname string) {
	// If the length of the hostname parameter is greater than 10,
	// use slicing syntax to shorten to 10 characters.
	if len(hostname) > 10 {
		hostname = hostname[:10]
	}

	// Assign the result to the `hostname` field of receiver `d`
	d.hostname = hostname
}

func main() {

	// instantiate a vlan type using the literal syntax.
	//
	// You can populate every field with a value, or you can leave it out, and the
	// field will be set to the zero-value for that field's type.
	myVlan := vlan{
		id:   5,
		name: "VLAN_5",
	}
	fmt.Println(myVlan)

	// We can also set these fields after instantiation
	myVlan.id = 6
	myVlan.name = "VLAN_6"

	// Methods are defined on a struct instance, so we must first instantiate `device`
	// as a new variable `myDevice`.
	myDevice := device{hostname: "r1"}
	// While functions are called from a package (i.e. `fmt`), methods are called
	// from a instance of a struct, which we created above.
	//
	// Note that there's no need for this method to have a `hostname` parameter;
	// since the receiver is passed implicitly, the method already has access to
	// this receiver's `hostname` field.
	myDevice.printHostname() // output: "r1"

	myDevice.setHostnameValueReceiver("r2")
	myDevice.printHostname() // output: "r1" ??

	// Since the `setHostname()` method is declared with a pointer receiver,
	// it will mutate the `hostname` field in the original instance, represented here
	// by the variable `myDevice`.
	myDevice.setHostname("r2")
	myDevice.printHostname() // output: "r2"

	// vlanMap := map[uint]string{
	// 	10: "VLAN_10",
	// 	20: "VLAN_20",
	// 	30: "VLAN_30",
	// }

}
