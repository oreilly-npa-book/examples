package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channelsExample()
}

func goroutineExample() {

	// This call to time.Sleep() will "block", meaning it will halt execution
	// of the rest of this program until the timer expires.
	time.Sleep(1 * time.Second)

	// This call to time.Sleep() will not block execution of this program,
	// because it is launched as a goroutine. The timer will still count down,
	// but this will happen in a separate lightweight thread, so the following
	// instruction(s) will execute immediately.
	go time.Sleep(10 * time.Second)

	// We will see this immediately - not after 10 seconds.
	fmt.Println("Program finished!")
}

func waitGroupExample() {

	var wg sync.WaitGroup
	var numGoroutines = 5

	// This is where we tell our wait group how many goroutines to wait for
	wg.Add(numGoroutines)

	// This allows us to launch all our goroutines in quick succession
	for i := 1; i <= numGoroutines; i++ {

		go func(i int) {

			// This will decrement (subtract) our wait group by a value of 1.
			// Remember, defer statements run at the END of a function.
			defer wg.Done()

			fmt.Printf("Goroutine started with duration of %d seconds\n", i)
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Printf("%d second goroutine finished!\n", i)
		}(i)
	}

	// This will block until all of the goroutines have finished
	wg.Wait()

	fmt.Println("Program finished!")
}

func channelsExample() {

	// Channels have a type (in this case float32), as well as a length. By omitting the length
	// parameter to make(), we're creating an "unbuffered channel", which has a length of 0.
	fChan := make(chan float32)

	// getDeviceCPU simulates an API call to a network device to get the current CPU utilization
	getDeviceCPU := func() float32 {
		time.Sleep(250 * time.Millisecond)
		return rand.Float32()
	}

	// Here, we're doing a fairly common task of retrieving a device's CPU via an API call.
	// We want to do ongoing monitoring (and keep sending values into the channel) so we'll do this
	// in an infinite loop.
	go func(iChan chan float32) {
		for {
			cpu := getDeviceCPU()

			// For values less than 0.8, we don't really care. But if it's higher, let's notify the
			// main goroutine by sending this value on the channel.
			if cpu >= 0.8 {

				// The "send" syntax places the channel on the left side of the `<-` operator
				// Remember, this will block execution of the goroutine until we receive a value
				// from this channel in the main goroutine.
				iChan <- cpu
			}
		}
	}(fChan)

	// This is an infinite loop so that our program continually receives values off of the channel.
	for {

		// Remember, channels allow for synchronization of goroutines, as well as conveying values.
		// When receiving a value here, we know that the goroutine is sending a value at the same time.
		// Remember, this will block execution of this main goroutine until the goroutine we launched
		// earlier sends a value into the channel.
		//
		// The "receive" syntax places the channel on the right side of the `<-` operator.
		fmt.Println(<-fChan)
	}
}

func mutexExample() {

	// We're declaring the map and the mutex here. Because they are contained in the outer scope, they can
	// be referenced directly by the goroutines we'll launch later on.
	var cpuMap = make(map[string]float32)
	var cpuMapMut = sync.Mutex{}

	// getDeviceCPU simulates an API call to a network device to get the current CPU utilization
	getDeviceCPU := func() float32 {
		return rand.Float32()
	}

	// monitorFunc is the function that we'll eventually launch as a goroutine, and it contains the infinite for loop
	// as well as the updates to the cpuMap, including the mutex Lock and Unlock operations
	monitorFunc := func(hostname string) {
		for {
			cpu := getDeviceCPU()

			// This call will block execution if another goroutine already has a lock. Only when we're able
			// to successfully acquire a lock in **this** goroutine will execution continue. This is how
			// we can safely write to a map from multiple concurrent goroutines.
			cpuMapMut.Lock()

			// Now that we have a Lock on the mutex, we can write to the map safely. Without the mutex (or
			// some other similar tool which offers the same kind of synchronization), our program might crash
			// when multiple goroutines try to access the map at the same time.
			cpuMap[hostname] = cpu

			// Don't forget to unlock the mutex when we're done, so that other goroutines can use it!
			// Sometimes you'll see `defer` used to call `Unlock()` automatically at the end of the function.
			cpuMapMut.Unlock()
		}
	}

	// Launch three different goroutines, one for each device.
	go monitorFunc("sw01")
	go monitorFunc("sw02")
	go monitorFunc("sw03")

	// Repeatedly print the contents of the map to the scree
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("cpuMap: %v\n", cpuMap)
	}

}
