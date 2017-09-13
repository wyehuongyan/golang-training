package main

import (
	"fmt"
)

func main() {
	// using make
	var myGreeting1 = make(map[string]string) // is not nil
	myGreeting1["Tim"] = "Good morning."
	myGreeting1["Jenny"] = "Bonjour."

	fmt.Println(myGreeting1)

	// nil map
	var myGreeting2 map[string]string

	if myGreeting2 == nil {
		fmt.Println("nil map")
	} else {
		fmt.Println("myGreeting2 is not nil")
	}

	// composite literal
	var myGreeting3 = map[string]string{"Huong": "Ohayo"}
	// myGreeting3 := map[string]string{"Huong": "Ohayo"} // this is the shorthand using :=

	if myGreeting3 == nil {
		fmt.Println("nil map")
	} else {
		fmt.Println("myGreeting3 is not nil")
	}

	myGreeting3["Yan"] = "Good morning."
	myGreeting3["Wye"] = "Bonjour."

	fmt.Println(myGreeting3)

	// "comma ok" idiom, does value exist?
	if val, exists := myGreeting3["He"]; exists {
		fmt.Println("Does exist!")
		fmt.Println(val, exists)
	} else {
		fmt.Println("Does not exist!")
		fmt.Println(val, exists)
	}

	rangeLoop()
}

func rangeLoop() {
	myGreeting := map[int]string{
		0: "Ni Hao",
		1: "Good morning",
		2: "Ohayo",
		3: "Lei Hou",
	}

	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
}
