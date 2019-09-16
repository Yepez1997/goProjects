package main

import (
	"fmt"
	"sync"
)

// fan in fan out patterns
// multiple functions can read from the same channels until that channel is closed
// distribute work amongst a group of workers to parallelize CPU use and IO
// a funcion can read from multiple inputs ad proceed untila all are closed by multiplexing the input channels onto a single channel thats clsoed when all the inputs are closed
// this is known as fan in

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// func gen(nums ...int) <-chan int {
// 	out := make(chan int, len(nums))
// 	for _, n := range nums {
// 		out <- n
// 	}
// 	close(out)
// 	return out
// }

func sq(in <-chan int) <-chan int {
	// received the square inputs here
	out := make(chan int)
	// ideally add a waithgrou here
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// merge function converst a lsit of channels to a single channel by starting a go routine for each
// inbound channel that copies tha values to the sole outbound channel
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	// start an output go routine for each input channel in cs
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// add as many go routines as the size of the channel
	wg.Add(len(cs))

	// for each input send to output channel
	for _, c := range cs {
		go output(c)
	}

	// start a go routine to close out all once the go routines are done
	go func() {
		wg.Wait()
		close(out)
	}()
	return out

}

func main() {

	in := gen(2, 3)
	// distribute the sq work accorss two go routines that broth read from in
	c1 := sq(in)
	c2 := sq(in)

	// consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}

}

// misc
// pattern to channels
// stages close theri outbound channels whe all the send operations are done
// stages keep receiving values from inbound channels until those channels are closed
