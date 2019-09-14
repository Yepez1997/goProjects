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
	// place a lock in the critical section
	mutex.Lock()
	fmt.Println(i)
	total += i
	fmt.Printf("Current total after summing %d: %d\n", i, total)
	time.Sleep(1 * time.Second)
	// unlock the critical section
	mutex.Unlock()
	// go routines finished so decrement by calling done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup // adding a wait group so that all go routines finish

	numberGoRoutines := 0
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			// add to the wait group to keep track ofthe number
			wg.Add(1)
			numberGoRoutines++
			go randomInt(i, &wg)
		}
		//time.Sleep(10 * time.Second)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			numberGoRoutines++
			go randomInt(i, &wg)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Done with all %d go routines\n", numberGoRoutines)
	fmt.Printf("Sum of all go routines -> %d\n", total)
}
