package main

import (
	"fmt"
	// "time"

	"github.com/scrapli/scrapligo/platform"
	"github.com/scrapli/scrapligo/driver/options"
)

func main() {
	p, err := platform.NewPlatform(
		"arista_eos",
        "eos-spine1",
		options.WithPort(22),
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("ntc"),
		options.WithAuthPassword("ntc123"),
	)
	if err != nil {
		fmt.Printf("failed to create driver; error: %+v\n", err)
		return
	}

    d, err := p.GetNetworkDriver()
    if err != nil {
        fmt.Printf("failed to fetch network driver from the platform; error: %+v\n", err)

        return
    }

	// returns err only
	err = d.Open()
	if err != nil {
		fmt.Printf("failed to open driver; error: %+v\n", err)
		return
	}
	defer d.Close()

	// Confirm we can get prompt
	prompt, err := d.GetPrompt()
	if err != nil {
		fmt.Printf("failed to get prompt; error: %+v\n", err)
		return
	}
	fmt.Printf("found prompt: %s\n\n\n", prompt)

	// Send a single command
	r, err := d.SendCommand("show version | i uptime")
	if err != nil {
		fmt.Printf("failed to send command; error: %+v\n", err)
		return
	}
	fmt.Printf(
		"sent command '%s', output received (SendCommand):\n %s\n\n\n",
		r.Input,
		r.Result,
	)

	// Single command passed from variable
	cmd := "show version"
	r, err = d.SendCommand(cmd)
	if err != nil {
		fmt.Printf("failed to send command; error: %+v\n", err)
		return
	}
	fmt.Printf(
		"output received (SendCommand):\n %s\n\n\n",
		r.Result,
	)

	// Multiple commands from a string slice
	cmds := []string{"show uptime", "show ip route"}
	rs, err := d.SendCommands(cmds)
	if err != nil {
		fmt.Printf("failed to send commands; error: %+v\n", err)
		return
	}
	for _, r := range rs.Responses {
		fmt.Println("------------------")
		fmt.Printf("Output for cmd '%v'\n", r.Input)
		fmt.Println("------------------")
		fmt.Printf("%v\n\n", r.Result)
	}

	// force timeout
	// ms, _ := time.ParseDuration("1ms")
	// r, err = d.SendCommand("show version | i uptime", options.WithSendTimeoutOps(ms))
	// if err != nil {
	//	panic(err)
	// }
	// fmt.Printf(
	// 	"sent command '%s', output received (SendCommand):\n %s\n\n\n",
	//	r.Input,
	//	r.Result,
	//)
}
