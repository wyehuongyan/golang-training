package main

import "fmt"

func main() {
	if true || false {
		fmt.Println("This ran")
	}

	//var xs = []int{1, 2, 3}
	xs := []int{1, 2, 3}

	if xs[0] == 1 && xs[2] == 3 {
		fmt.Println(xs)
	}
}
