// or done channel in go
// combine one or more done channels into a single done channel that closes if any
// of its component cloeses
package main

import "time"

import "fmt"

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		fmt.Printf("%v here sig\n", after)
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		fmt.Printf("%d len \n", len(channels))
		switch len(channels) {
		case 0:
			fmt.Println("HERE 2")
			return nil
		case 1:
			fmt.Println("HERE")
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)

			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(5*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
