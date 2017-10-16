package main

import "fmt"
import "sync"

/*
Execute 100 factorial computations concurrently and in parallel
*/

func main() {
	nums := make([]int, 0)

	for i := 0; i < 10; i++ {
		for j := 3; j < 13; j++ {
			nums = append(nums, j)
		}
	}

	// using channels
	d := factorial(nums...)
	for i := 0; i < len(nums); i++ {
		fmt.Println("Factorial: ", <-d)
	}
}

func factorial(nums ...int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, n := range nums {
		go func(n int) {
			total := 1
			for i := n; i > 0; i-- {
				total *= i
			}
			out <- total
			wg.Done()
		}(n)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
