package main

import (
	"fmt"
)

func main() {
	// not using channels
	//f := factorial(4)
	//fmt.Println(f)

	// using channels
	c := totalizer(decrementor(4))
	fmt.Println("Total: ", <-c)

	d := factorial(4)
	for i := range d {
		fmt.Println("Factorial: ", i)
	}
}

/*
func factorial(n int) int {
	total := 1

	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
*/

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}

// one goroutine to decrement n
// one goroutine to receive decrements, and calculate total
// my style ;)
func decrementor(n int) chan int {
	out := make(chan int)

	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()

	return out
}

func totalizer(c chan int) chan int {
	out := make(chan int)

	go func() {
		total := 1
		for i := range c {
			total *= i
		}
		out <- total
		close(out)
	}()

	return out
}
