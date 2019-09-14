package main

import (
	"fmt"
	"sync"
	"time"
)

// random concurrency stuff to test out
// channels
// anonymous functions
var total = 0
var mutex = sync.Mutex{}

func randomInt(i int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println(i)
	total += i
	fmt.Printf("Current total after summing %d: %d\n", i, total)
	time.Sleep(1 * time.Second)
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup // adding a wait group so that all go routines finish

	numberGoRoutines := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		numberGoRoutines++
		go randomInt(i, &wg)
	}
	wg.Wait()
	fmt.Printf("Done with all %d go routines\n", numberGoRoutines)
	fmt.Printf("Sum of all go routines -> %d\n", total)
}
