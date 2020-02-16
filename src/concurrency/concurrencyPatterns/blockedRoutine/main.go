// reverse from go routine leak folder
// go routine blocked on attemping to write a value to the channel
package main

import "fmt"

import "math/rand"

import "time"

func main() {
	newRandomStream := func(done <-chan interface{}) <-chan int {
		randStream := make(chan int)
		go func() {
			defer fmt.Println("newRandStreamClosed ...")
			defer close(randStream)
			for {
				select {
				case randStream <- rand.Int():
				case <-done:
					return
				}
			}
		}()
		return randStream
	}

	done := make(chan interface{})
	randStream := newRandomStream(done)
	fmt.Println("three random ints")
	for i := 1; i <= 3; i++ {
		// receive ?
		fmt.Printf("%d: %d\n", i, <-randStream)
	}
	close(done)

	time.Sleep(1 * time.Second)
}
