package main

import (
	"fmt"
	"time"
)

func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Go Routine")

	// add go routines
	go compute(10)
	go compute(10)

	// wait til finished
	var input string
	fmt.Scanln(&input)
}
