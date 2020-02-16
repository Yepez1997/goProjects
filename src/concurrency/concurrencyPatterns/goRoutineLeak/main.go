// preventing go routine leaks
package main

import "fmt"

import "time"

// cleaning go routines
// add the done channel so that the parent go routine can signal the child to stop

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("do work exited")
			// close the channel
			defer close(terminated)
			// for s := range strings {
			// 	// do something interesting
			// 	fmt.Println(s)
			// }
			for {
				select {
				case s := <-strings:
					// also do something interesting .... ?
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling do work go routine")
		close(done)
	}()

	<-terminated

	fmt.Println("Done ....")
}
