// lexical confinement
// lexical confinement hapepns at the compiler level
package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		// lexical scope
		// limits the scope of chan owner
		// confines the write aspect
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	// receive a read only copy of the int instance
	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Recieved: %d\n", result)
		}
		fmt.Println("Done receiving ...")
	}

	// provide a read only view
	results := chanOwner()
	consumer(results)
}
