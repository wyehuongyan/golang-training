package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // 2 means wait for 2 go routines to finish (wait group counter)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done() // decrements waitgroup counter by 1
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done() // decrements waitgroup counter by 1
	}()

	go func() {
		wg.Wait() // wait blocks until the waitgroup counter is 0
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
