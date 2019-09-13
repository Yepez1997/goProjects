package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my go routine")
	// indicate the go routine has finished
	waitgroup.Done()
}

// simple waitgroup
func main() {
	fmt.Println("Hello World")
	var waitgroup sync.WaitGroup
	// call .add before attemtping to execute go routine
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	// do not go futher until the go routine has finished executing
	waitgroup.Wait()
	// once finished print this
	fmt.Println("Finished Execution")

}
