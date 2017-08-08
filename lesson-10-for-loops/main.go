package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, " - ", j)
		}
	}

	k := 0
	for {
		fmt.Println(k)
		k++
		if k > 10 {
			break
		}
	}

	counter := true
	for counter == true {
		fmt.Println("true")

		counter = false
	}
}
