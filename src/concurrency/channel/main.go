package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	// send to values
	values <- value
}

// channels in go
// channels are pipes that link go routines
func main() {
	fmt.Println("Go Channel")
	values := make(chan int)
	defer close(values)
	go CalculateValue(values)
	// receive from values
	value := <-values
	fmt.Println(value)
}
