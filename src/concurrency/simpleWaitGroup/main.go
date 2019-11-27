package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my go routine")
	waitgroup.Done()
}

func myFunc2(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside second go routine")
	waitgroup.Done()
}

func forFunc(waitgroup *sync.WaitGroup, n int) {
	fmt.Print("Inside for .... \n")
	for i := 0; i < n; i++ {

	}
	waitgroup.Done()
}

// simple waitgroup
func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	// call .add before attemtping to execute go routine
	waitgroup.Add(2)
	go myFunc2(&waitgroup)
	go myFunc(&waitgroup)
	n := 5

	for i := 0; i < n; i++ {
		waitgroup.Add(1)
		// these happpen concurrently
		go forFunc(&waitgroup, n)
	}

	// do not go futher until the go routine has finished executing
	waitgroup.Wait()
	// once finished print this
	fmt.Println("Finished Execution")

}
