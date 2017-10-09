package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // unbuffered channel (no limit on size)

	// blocking nature of a unbuffered channels between two goroutines
	// runners in a relay race
	go func() {
		for i := 0; i < 10; i++ {
			c <- i // put a number into the channel (code stops here until takes the number off, relay race)
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // receives the number off from the channel (code stops until new value comes into channel)
		}
	}()

	time.Sleep(time.Second) // wait a sec for the above goroutines to execute
}
