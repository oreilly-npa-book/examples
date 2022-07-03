package main

import (
	"fmt"
)

func main() {

	r1 := Router{hostname: "r1"}
	printHostname(r1)

	s1 := Switch{hostname: "s1"}
	printHostname(s1)

	f1 := Firewall{hostname: "f1"}
	printHostname(f1)

	///

	rtr := Router{hostname: "rtr1-dc1"}

	// Fails to compile!
	//
	// ./10-interfaces.go:23:23: cannot use rtr (variable of type Router) as
	//      type Trimmable in argument to printHostnameTrimmed:
	// Router does not implement Trimmable (TrimHostname method has pointer receiver)
	// printHostnameTrimmed(rtr, 4)

	// We can create a pointer to a value using the ampersand symbol. Where `rtr` was
	// type `Router`, `rtrPointer` is type `*Router` (pointer of type `Router`)
	rtrPointer := &rtr

	// This works, because we're passing a pointer (*Router) to printHostnameTrimmed()
	// rather than a value. This means the method set now includes the method required
	// to satisfy Trimmable
	printHostnameTrimmed(rtrPointer, 4)

	// We could also skip defining a separate variable and do this all in one step
	printHostnameTrimmed(&rtr, 4)

}

// This interface type describes any concrete type
// which implments a `GetHostname()` method that has
// no parameters, and a single `string` return type.
type Hostnamer interface {
	GetHostname() string
}

type TrimmableEmbedded interface {
	// The Hostnamer interface is embedded here, which means all of the methods from that
	// interface are included alongside the other methods listed here.
	Hostnamer
	TrimHostname(int)
}

type Router struct {
	hostname string
	vrfs     []string
}

// This method allows the Router type to satisfy
// the Hostnamer interface
func (r Router) GetHostname() string {
	return r.hostname
}

type Switch struct {
	hostname string
	vlans    []int
}

func (s Switch) GetHostname() string {

	// There's no rule that says we **have** to return r.hostname directly.
	// What we do inside the method doesn't affect whether or not it implements
	// the Hostnamer interface. We can give the hostname a prefix of `switch-`!
	return fmt.Sprintf("switch-%s", s.hostname)
}

type Firewall struct {
	hostname string
	zones    []string
}

func (f Firewall) GetHostname() string {
	return fmt.Sprintf("firewall-%s", f.hostname)
}

// This function takes a concrete type (the "Router" struct)
// and therefore, no other type can be used when calling this
// function.
func printHostnameConcrete(device Router) {
	fmt.Printf("The hostname is %s\n", device.hostname)
}

// This function uses the "Hostnamer" interface for the "device" parameter,
// so any type which implements that interface can be used
func printHostname(device Hostnamer) {
	fmt.Printf("The hostname is %s\n", device.GetHostname())
}

type Trimmable interface {
	// Interfaces can specify more than one method
	TrimHostname(int)
	GetHostname() string
}

// Remember, when mutating the fields of the receiver, we usually want
// to use a pointer receiver. Otherwise we'll just mutate a copy of the
// receiver value.
func (r *Router) TrimHostname(length int) {
	// This syntax trims the string so that it's no longer than
	// the "length" parameter. Of course, if it's already shorter,
	// we don't need to do anything.
	if len(r.hostname) > length {
		r.hostname = r.hostname[:length]
	}
}

func printHostnameTrimmed(device Trimmable, trimLength int) {
	// Trimmable requires the TrimHostname method to be defined, so we
	// know we can use it here
	device.TrimHostname(trimLength)

	// Trimmable also uses the GetHostname method so we can use this
	// to retrieve the result after we've trimmed it.
	fmt.Printf("The device hostname trimmed to %d characters is %s\n", trimLength, device.GetHostname())
}
