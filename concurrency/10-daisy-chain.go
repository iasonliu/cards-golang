package main

import "fmt"

//  Daisy chaining goroutines... all routines at once, so if one is fulfilled
//  everything after should also have their blocking commitment (input) fulfilled

// takes tow int channels, stores right val (+1) into left
func fw(left, right chan int) {
	// after 1st right read, locks until left read
	left <- 1 + <-right
}

func main() {
	const n = 100000
	// construct an array of n+1 in channels
	var channels [n + 1]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}
	// wire n goroutines in a chain
	for i := 0; i < n; i++ {
		go fw(channels[i], channels[i+1])
	}

	// insert a value into right-hand end
	go func(c chan<- int) {
		c <- 1
	}(channels[n])

	// get value form left-hand end
	fmt.Println(<-channels[0])

}
