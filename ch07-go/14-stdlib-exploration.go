package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/netip"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {
	workingWithStrings()
}

func workingWithStrings() {
	exampleString := `
	Hello network automators! Welcome to Network Programmability and Automation.
	`

	// HasPrefix, HasSuffix allow you to look specifically at the beginning or end of
	// the string, respectively
	doesContain := strings.Contains(exampleString, "Automation")
	fmt.Println(doesContain) // output: true

	// strings.Index() returns the index (location within the string) of the first
	// instance of the substring.
	substringIndex := strings.Index(exampleString, "Welcome")
	fmt.Println(substringIndex) // output: 27

	// strings.Split() will create a slice of strings ([]string) from the input string
	// based on a provided delimiter. The below example uses a space as a delimiter,
	// which will result in each word of the input string being placed in its own
	// slice element.
	//
	// You can perform the reverse of this operation using strings.Join(), which
	// creates a single string from a []string, joined with a delimiter of your choice.
	strSplit := strings.Split(exampleString, " ")
	fmt.Println(strSplit[4]) // output: "Welcome"

	// strings.TrimSpace() is a super handy function for easily removing extra spaces
	// at the beginnin or end of a string.
	//
	// There are plenty of other trim functions in `strings`, each with their own
	// specialized use cases, including Trim, TrimPrefix/TrimSuffix, and
	// TrimLeft/TrimRight
	strTrimmed := strings.TrimSpace("    Automation!    ")
	fmt.Println(strTrimmed)      // output: "Automation!"
	fmt.Println(len(strTrimmed)) // output: "11"

	// strings.ReplaceAll() can be used to replace all instances of a given substring
	// with another string of your choice.
	//
	// If you only want to replace a limited number of instances, you can use
	// strings.Replace().
	strReplaced := strings.ReplaceAll(exampleString, "network", "gopher")
	fmt.Println(strReplaced)
	// output: "Hello gopher automators! Welcome to Network Programmability and Automation."

	// strconv.Atoi converts a string to an integer. It returns the integer
	// value but also an err, as the parse might fail due to integer overflow,
	// non-integers, etc.
	i, err := strconv.Atoi("-42")
	if err != nil {
		fmt.Printf("Unable to convert string to integer: %s\n", err)
	} else {
		fmt.Printf("Parsed integer is %d\n", i)
	}

	// strconv.ItoA performs the reverse, converting an integer to a string.
	// This cannot fail, so we only see one return type from this function.
	i42 := strconv.Itoa(42)
	fmt.Printf("i42 as a string is %s\n", i42)

	// outputStr is a large multi-line string we can use to parse using the
	// `regexp` package.
	outputStr := `
    eth0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
        inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
        ether 02:12:2a:24:5b:98  txqueuelen 0  (Ethernet)
	`

	// This regular expression matches MAC addresses.
	// regexp.Compile() returns a *regexp.Regexp, which we can use
	// for later tasks. All following tasks are done using methods
	// of this returned instance `re`.
	//
	// This is a common step in many implementations of regular expressions,
	// including in languages outside Go. This helps ensure that we have a
	// valid expression before we continue.
	re, err := regexp.Compile(`([0-9a-f]{2}:){5}[0-9a-f]{2}`)
	if err != nil {
		panic(err)
	}

	// We can use the `MatchString()` method of the returned instance `re`
	// to get a basic boolean true/false to indicate if there is any substring
	// in outputStr that matches our regular expression.
	fmt.Println(re.MatchString(outputStr)) // output: true

	// FindString() goes a step further and returns the first specific substring that
	// matched our expression.
	//
	// Note that other methods like FindAllString() and FindAllStringIndex() can be used
	// to find all instances which match, returning a slice of strings ([]string) which
	// can be inspected afterwards, or perhaps even iterated over.
	fmt.Println(re.FindString(outputStr)) // output: 02:12:2a:24:5b:98

	// ReplaceAllString() allows you to replace all instances that match the
	// expression with a given string literal. In this case we're overwriting the
	// MAC address with all 0s, leaving the rest untouched.
	//
	// output:
	//
	// eth0: flags=4099<UP,BROADCAST,MULTICAST>  mtu 1500
	//     inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
	//     ether 00:00:00:00:00:00  txqueuelen 0  (Ethernet)
	fmt.Println(re.ReplaceAllString(outputStr, "00:00:00:00:00:00"))

}

func dataSerialization() {

	// An important thing to keep in mind for the `encoding` package (and generally
	// any package which performs serialization/deserialization) is that it will only
	// work with exported (starts with capital letter) types and fields.
	type NetworkInterface struct {

		// You may wonder about this strange string after this field. This is called
		// a "struct tag", and while not required, it is extremely common to see these
		// for structs that will be used for serialization/deserialization purposes,
		// such as to/from JSON or XML.
		//
		// Generally, struct tags are just metadata - they have no implicit purpose on
		// their own. However, both the `xml` and `json` package can use these if present
		// to specify a field name which is different from the actual struct field's name
		Name  string `xml:"name" json:"name"`
		Speed int
	}

	type Device struct {
		Hostname   string
		Interfaces []NetworkInterface
	}

	r1 := Device{
		Hostname: "r1",
		Interfaces: []NetworkInterface{
			{
				Name:  "eth0",
				Speed: 1000,
			},
		},
	}

	jsonOut, err := json.Marshal(&r1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonOut))
	// output:  {"Hostname":"r1","Interfaces":[{"name":"eth0","Speed":1000}]}

	xmlOut, err := xml.Marshal(&r1)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xmlOut))
	// output:  <Device><Hostname>r1</Hostname><Interfaces>
	//             <name>eth0</name><Speed>1000</Speed></Interfaces></Device>

}

// example 3
func workingWithNetworking() {

	// The encoding/http package includes high-level functions like Get() for
	// performing requests easily with some sensible defaults.
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// We can use what we learned in the previous example with the "encoding" package
	// to unmarshal the raw JSON string into a struct type
	ipifyResponse := struct {
		IP string `json:"Ip"`
	}{}
	err = json.Unmarshal(body, &ipifyResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(ipifyResponse.IP)

	// Sometimes you need a bit more control than these high level functions offer. For
	// instance, you may need to send specific HTTP headers with your request. This
	// requires that you create your own Client and Request.
	client := &http.Client{}
	// Note that the method (GET) is defined here, as is the URL.
	req, err := http.NewRequest("GET", "https://api.ipify.org?format=json", nil)
	if err != nil {
		panic(err)
	}
	// Headers are defined on the request object
	req.Header.Add("My-Header", `foo`)
	// Once prepared, the request is passed as a parameter to the function Do(), a method
	// on the client we created earlier.
	resp, err = client.Do(req)

	// The net.IP type is used to represent IP addresses. While it includes many
	// helpful methods, at its core it is really just a slice of bytes, so we can
	// initialize a net.IP instance by constructing this byte slice ourselves
	var ipFromByteSlice net.IP = []byte{192, 168, 0, 1}
	fmt.Println(ipFromByteSlice)

	// For convenience, net.ParseIP() allows us to construct IP instances
	// with a string as input. This is a much more common way of constructing
	// net.IP instances.
	//
	// You'll notice that either IPv4 or IPv6 addresses can be passed here.
	// This is due to the flexibility of the byte slice representation, with a
	// length suited to the address being represented.
	addrOne := net.ParseIP("192.168.0.1")
	addrTwo := net.ParseIP("2001:db8::1")
	fmt.Println(addrOne)
	fmt.Println(addrTwo)

	// net.IPNet is used to represent a network/subnet. It is defined by two fields:
	// an IP (net.IP) and a Mask.
	//
	// Like net.IP, net.IPNet is v4/v6 agnostic.
	network := net.IPNet{
		IP: net.ParseIP("192.168.0.0"),
		// Here, we're defining a bitmask of 24 bits long, with a total size of 32 bits.
		// In other words, this is the v4 subnet mask 255.255.255.0
		Mask: net.CIDRMask(24, 32),
	}
	// We can then use the Contains() method on this type to very easily determine if
	// a given IP address is a member of this network.
	fmt.Println(network.Contains(addrOne)) // output: true
	fmt.Println(network.Contains(addrTwo)) // output: false

	// ParseAddr allows us to parse an IP address (v4 or v6) from a string. Once we have the resulting
	// netip.Addr type, we can call helpful methods like IsGlobalUnicast() or IsLoopback() to quickly
	// identify properties of the address we parsed.
	ipv6, err := netip.ParseAddr("2001:db8::1")
	if err != nil {
		panic(err)
	}
	fmt.Println(ipv6.IsGlobalUnicast()) // output: true

	// ParseAddr does work for IPv4 addresses as well, but an alternative is the `AddrFrom4()` function,
	// which allows us to pass the address as a 4-length byte array, removing the need for
	// error handling (no parsing is being done here).
	fmt.Println(netip.AddrFrom4([4]byte{127, 0, 0, 1}).IsLoopback()) // output: true

	// We can parse entire prefixes from a string using `ParsePrefix()`
	prefixString := "192.168.0.0/24"
	prefix, err := netip.ParsePrefix(prefixString)
	if err != nil {
		panic(err)
	}
	fmt.Println(prefix)
}

func workingWithTime() {
	// the time.Time type is a singular point in time. One of the most common
	// ways of getting this is via the Now() function, which returns the
	// current time.
	now := time.Now()
	fmt.Println(now)

	// However, you can create an instance of time.Time representing any arbitrary
	// date/time:
	moonLanding := time.Date(1969, time.July, 20, 20, 17, 45, 0, time.UTC)

	// time.Duration is actually a type alias for int64, and it's used to represent
	// a time duration in nanoseconds. The below example is equivalent to one second
	// of duration.
	var oneSecond time.Duration = 1000000000
	fmt.Println(oneSecond)

	// However, the time package also includes convenient constants which make
	// it easier to represent durations in a more readable way:
	tenSeconds := 10 * time.Second

	// time.Since() is a common way to derive a Duration between some event in the past,
	// and the current time
	fmt.Println(time.Since(moonLanding))

	// And there's the ever-useful Sleep() function, which - you guessed it - sleeps
	// the current goroutine for the specified Duration.
	time.Sleep(tenSeconds)

}

func workingWithOs() {

	// Reading a text file is easy with os.ReadFile
	dat, err := os.ReadFile("sampleconfig.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dat))

	// Next, let's write a file. Here we're first marshaling a struct into JSON
	// so we can write the result to the file.
	jsonOut, err := json.Marshal(struct {
		Hostname   string
		Interfaces []string
	}{
		"sw01",
		[]string{"eth0", "eth1", "eth2"},
	})
	if err != nil {
		panic(err)
	}
	// Just like ReadFile returns a []byte value, so does WriteFile require
	// this type as an argument. Fortunately that's exactly what json.Marshal returns
	err = os.WriteFile("sampleconfig.json", jsonOut, 0644)
	if err != nil {
		// Here, instead of calling panic(), we can use os.Exit to more gracefully exit our program,
		// while returning an error code to the operating system.
		fmt.Printf("Unable to write file: %s\n", err)
		os.Exit(1)
	}

	// We can use a combination of packages like os, os/signal, and syscall
	// to handle incoming signals from the operating system. This allows us to
	// more gracefully handle these signals. Here, we're creating a channel
	// of type os.Signal
	sigs := make(chan os.Signal, 1)
	// We then pass this channel into signal.Notify(), along with a list of signals
	// we wish to handle.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Let's launch a goroutine to simulate doing some actual work
	go func() {
		for {
			fmt.Println("Doing some work...")
			time.Sleep(1 * time.Second)
		}
	}()
	// If the operating system sends any of the above-listed signals to our application,
	// this channel will receive this, and the code below will execute. However, until then
	// it will block, as it is an unbuffered channel.
	<-sigs
	// Here, we added a simple print statement, but we could add any logic we wanted,
	// to make sure we take care of any clean-up tasks before our program shuts down.
	fmt.Println("exiting")
	os.Exit(0)
}
