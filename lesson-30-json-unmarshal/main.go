package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.First) // prints zero-value (empty)
	fmt.Println(p1.Last)  // prints zero-value (empty)
	fmt.Println(p1.Age)   // prints zero-value (0)
	fmt.Printf("%T\n", p1)

	// notice the round brackets, this is conversion
	bs := []byte(`{"First":"James", "Last":"Bond", "wisdom score":20}`) // conversion: string to a slice of bytes (this how json is represented)
	json.Unmarshal(bs, &p1)

	fmt.Println("------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T\n", p1)
}
