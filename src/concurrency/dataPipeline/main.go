package main

import "fmt"

// dataPipelines -s sqauring numbers in a data pipeline
// informal definition of a pipeline (below)
// pipeline is a series of stages connected by channels
// where eacb stage is a group of go routines running in the same function
// at each stage go routines
// receive values from upstream methods via inbound channels
// perform some function on that data usually producing new values
// send values downsteam vai outbound channels

// each stage has any number of inbound and outboud channels - except the first and last
// first stage is referred to the source / producer
// last stage is the sink or consumer

// first stage of the data pipeline
// converts a list of integers to a channel and closes the channel once all the values have been sent
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

// second stage of the data pipeline
// receives values from a channel and returns a channel that emits the square of each receieved integer
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

// data pipelines and cancellation in go
func main() {
	// set up the pipeline
	// c := gen(2, 3)
	// out := sq(c)
	// fmt.Println(<-out)
	// fmt.Println(<-out)

	// can combine as long as the same types exist for inbound and outbound channels
	for n := range sq(sq(gen(2, 3))) {
		fmt.Println(n)
	}

}
