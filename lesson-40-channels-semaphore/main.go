package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // you can only put something into the channel if someone is waiting to receive, line 41
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done // sits right here until done comes through so it can be pulled
		<-done // sits right here until done comes through so it can be pulled
		close(c)
	}()

	/*
		// wrong example, will prevent line 41 from running
		<-done
		<-done
		close(c)
	*/

	// while go routine is running somewhere else, code goes here
	// ready and waiting to receive off the unbuffered channel, back and forth
	// code halts here (do not need time.sleep)
	for n := range c {
		fmt.Println(n)
	}
}
