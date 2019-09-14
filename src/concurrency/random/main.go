package main

import (
	"fmt"
	"sync"
)

// TODO -- send the sum through a channel

// random concurrency stuff to test out
// channels
// anonymous functions
var total = 0
var mutex = sync.Mutex{}

func randomInt(i int, intStream chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	intStream <- i
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup // adding a wait group so that all go routines finish
	intStream := make(chan int, 20)

	numberGoRoutines := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			// add to the wait group to keep track ofthe number
			wg.Add(1)
			numberGoRoutines++
			go randomInt(i, intStream, &wg)
		}
		//time.Sleep(10 * time.Second)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			numberGoRoutines++
			go randomInt(i, intStream, &wg)
		}
		defer wg.Done()
	}()
	// close the int stream once all the go routines finished
	// close(intStream)
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for val := range intStream {
	// 		total += val
	// 	}
	// }()

	wg.Wait()
	close(intStream)
	for val := range intStream {
		total += val
	}

	fmt.Printf("Done with all %d go routines\n", numberGoRoutines)
	fmt.Printf("Sum of all go routines -> %d\n", total)
}
