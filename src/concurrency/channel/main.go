package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CalculateValue - send random value to the channel
func CalculateValue(values chan int) {
	min := 10
	max := 30
	value := rand.Intn(max-min) + min
	fmt.Println("Calculated Random Value: {}", value)
	// send to values
	time.Sleep(1500 * time.Millisecond)
	values <- value
	fmt.Println("Only executes after another goroutine performs a receive on the channel")
}

// channels in go
// channels are pipes that link go routines
func main() {
	fmt.Println("Go Channel")
	values := make(chan int)
	// good practice
	defer close(values)
	go CalculateValue(values)
	go CalculateValue(values)
	// receive from values
	value := <-values
	// act of sending and receiving from a channel are blocking
	// ie this program will not finish unless a value is recieved
	fmt.Println(value)
}
