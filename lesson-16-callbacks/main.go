package main

import "fmt"

func main() {
	// example 1
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	// example 2
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}

func visit(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	//var xs []int

	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T\n", callback)
	return xs
}
