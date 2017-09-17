package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Wye Huong", "Yan", 29, 123}

	json.NewEncoder(os.Stdout).Encode(p1)
}
