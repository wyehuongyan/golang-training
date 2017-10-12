/*
Informally, a pipeline is a series of stages connected by channels, where each
stage is a group of goroutines running the same function.

In each stage, the goroutines
- receive values from upstream via inbound channels
- perform some function on that data, usually producing new values
- send values downstream via outbound channels
*/

package main

import (
	"fmt"
)

func main() {
	// set up the pipeline
	c := gen(2, 3, 4)
	out := sq(c)

	// consume the output
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)

	// square them twice!
	d := sq(sq(gen(1, 2, 3, 4)))

	fmt.Print("Squared twice: ")
	for i := range d {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func gen(nums ...int) chan int {
	//fmt.Printf("%T\n", nums) // []int (integer slice)

	out := make(chan int)
	go func() {
		for _, n := range nums { // range on arrays and slices provides both the index and value for each entry.
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}
