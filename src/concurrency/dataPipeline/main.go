package main

// dataPipelines -s sqauring numbers in a data pipeline


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

// data pipelines and cancellation in go 
func main() {

}