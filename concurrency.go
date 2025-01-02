// Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides useful primitives, although you won't need them much in Go as there are other primitives. (See the next slide.)

// Channels
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
//
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and
//
//	// assign value to v.
//
// (The data flows in the direction of the arrow.)
//
// Like maps and slices, channels must be created before use:
//
// ch := make(chan int)

package main

import (
	"fmt"
	"time"
)

func concurrency() {
	go say("world")
	say("hello")
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	d := make(chan int, 10)
	go fibonacci(cap(d), d)
	for i := range d {
		fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
// ch := make(chan int, 100)
//
// A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

// ok is false if there are no more values to receive and the channel is closed.
//
// /Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
//
// Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop./ v, ok := <-ch
//

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// The select statement lets a goroutine wait on multiple communication operations.
//
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
//
//
// The default case in a select is run if no other case is ready.

// Use a default case to try a send or receive without blocking:
