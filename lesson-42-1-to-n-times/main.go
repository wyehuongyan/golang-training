package main

import (
	"fmt"
)

// one channel polled by n functions
func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	// one goroutine putting things onto the channel
	go func() {
		for i := 0; i < 100000; i++ {
			c <- i
		}
		close(c)
	}()

	// 10 goroutines pulling from the channel
	for i := 0; i < n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true // will push true to done channel n times
		}()
	}

	// hold on until all the above has finished executing
	for i := 0; i < n; i++ {
		<-done
	}
}
