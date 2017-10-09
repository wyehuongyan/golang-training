package main

import "fmt"

func main() {
	c := make(chan int)

	// go routine runs on its own
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c) // all done, close channel (cannot add more stuff to channel)
	}()

	// while go routine is running somewhere else, code goes here
	// ready and waiting to receive off the unbuffered channel, back and forth
	// code halts here (do not need time.sleep)
	for n := range c {
		fmt.Println(n)
	}
}
