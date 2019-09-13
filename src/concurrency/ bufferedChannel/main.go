package main

import (
	"fmt"
	"math/rand"
	"time"
)

// buffered channels

// CalculateValue - send random value to the channel
func CalculateValue(values chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	// send to values
	time.Sleep(1000 * time.Millisecond)
	values <- value
	fmt.Println("This executes regardless as it is non blocking")
}

// channels in go
// channels are pipes that link go routines
func main() {
	fmt.Println("Go Channel")
	values := make(chan int, 2)
	// good practice
	defer close(values)
	go CalculateValue(values)
	go CalculateValue(values)
	// receive from values
	value := <-values
	// act of sending and receiving from a channel are blocking
	// ie this program will not finish unless a value is recieved
	fmt.Println(value)
	time.Sleep(time.Second)
}
