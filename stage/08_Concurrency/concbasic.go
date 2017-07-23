package main

import (
	"fmt"
	"time"
)

// goroutines vs OS Threads
// Lighter weight (2k way smaller than threads)
// Go manages goroutines
// Less switching, other goroutines will run the same threads
// Safe communication

// GO's concurrency model use Actor model which is a Communicating Sequential Process

func main() {
	func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	func() {
		fmt.Println("What?")
	}()
}
