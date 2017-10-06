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

/*
Marshal => String
Encode => Stream

Encoder and decoder write struct to slice of a stream or read data from a slice of a stream
and convert it into a struct. Internally, it also implements marshal method.
The only difference is if you wanna play with string or bytes use marshal and
if any data you want to read or write to some writer interface, use encodes and decode.

https://stackoverflow.com/questions/21197239/decoding-json-in-golang-using-json-unmarshal-vs-json-newdecoder-decode
*/
