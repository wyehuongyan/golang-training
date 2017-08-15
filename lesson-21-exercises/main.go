package main

import "fmt"

func main() {
	// q1 + q2
	half := func(x int) (int, bool) {
		mod := x%2 == 0
		return x / 2, mod
	}

	fmt.Println(half(1))
	fmt.Println(half(2))

	// q3
	fmt.Println(largestNum(1, 2, 3, 6, 4, 5))

	// q4
	expr()

	//q5
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func largestNum(nums ...int) int {
	var largest int

	for _, v := range nums {
		if v > largest {
			largest = v
		}
	}

	return largest
}

func expr() {
	if (true && false) || (false && true) || !(false && false) {
		fmt.Println("expr")
	}
}

func foo(input ...int) {
	fmt.Println(input)
}
