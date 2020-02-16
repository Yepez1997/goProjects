// ad hoc confinement
// ad hoc confinement is set by convention, group of people, org, etc
// safe operations
// sync by sharing memory - mutex, semaphores
// sync by communicating - channels
// information is only avaible to one process

package main

import "fmt"

func main() {
	data := make([]int, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			handleData <- data[i]
		}
	}

	handleData := make(chan int)
	go loopData(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}
