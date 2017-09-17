package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string `json:"-"`            // similar to notExported
	Age         int    `json:"wisdom score"` // json object will use this key instead of Age
	notExported int    // lowercase fields are not exported in json
}

func main() {
	p := Person{"Wye Huong", "Yan", 29, 123}

	fmt.Println(p)

	bs, _ := json.Marshal(p) // convert to json format
	fmt.Println(bs)          // prints a slice of bytes []byte
	fmt.Println(string(bs))
}
