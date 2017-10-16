package main

import (
	"fmt"
)

var workerID int
var publisherID int

func main() {
	input := make(chan string)

	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)
}

// publisher pushes data into a channel
func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0

	for {
		dataID++
		fmt.Printf("publiher %d is pushing data\n", thisID)
		data := fmt.Sprintf("Data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

func workerProcess(in <-chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Printf("%d: waiting for input...\n", thisID)
		input := <-in
		fmt.Printf("%d: input is : %s\n", thisID, input)
	}
}

/*
There is no fan in here, only fan out. Why?

	definition of fan out:
	- Fan out means you have multiple functions reading from the same channel until that channel is closed.

	definition of fan in:
	- Fan in means a function can read from multiple inputs and proceed until all are closed by multiplexing the input
	- channels onto a single channel that's closed when all the inputs are closed.

Here we have publishers and workers reading from input, but we do not have many channels converging into one channel.
*/
