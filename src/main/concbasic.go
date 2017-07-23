package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

// goroutines vs OS Threads
// Lighter weight (2k way smaller than threads)
// Go manages goroutines
// Less switching, other goroutines will run the same threads
// Safe communication

// GO's concurrency model use Actor model which is a Communicating Sequential Process

func main() {
	runtime.GOMAXPROCS(3)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitGrp.Done()
		fmt.Println("What?")
	}()

	waitGrp.Wait()
	fmt.Println("Main Exit..")
}
