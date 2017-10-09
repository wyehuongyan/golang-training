package main

import "fmt"

func main() {
	// conversion: int to float 64
	var x = 12
	var y = 12.1230123
	fmt.Println(y + float64(x)) // widening conversion
	fmt.Println(int(y) + x)     // narrowing conversion

	// conversion: rune to string
	var a rune = 'a' // rune is an alias for int32
	var b int32 = 'b'

	fmt.Println(string(a))
	fmt.Println(string(b))

	// assertion .(type you are asserting it to be) for interfaces
	var name interface{} = "Sydney" // the concrete type that is implementing this empty interface is a string ("Sydney")

	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
