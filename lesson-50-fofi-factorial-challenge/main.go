package main

import (
	"fmt"
	"sync"
)

func main() {
	c := gen()

	// we have 10 workers to pull data off of c
	c1 := factorial(c)
	c2 := factorial(c)
	c3 := factorial(c)
	c4 := factorial(c)
	c5 := factorial(c)
	c6 := factorial(c)
	c7 := factorial(c)
	c8 := factorial(c)
	c9 := factorial(c)
	c10 := factorial(c)

	cs := []chan int{c1, c2, c3, c4, c5, c6, c7, c8, c9, c10}

	// merge the channels from c1 to c10 onto a single channel
	for n := range merge(cs) {
		fmt.Println("Factorial: ", n)
	}
}

// generate numbers
func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

// fan out
func factorial(in chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
		}
		close(out)
	}()

	return out
}

// fan in
func merge(cs []chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
