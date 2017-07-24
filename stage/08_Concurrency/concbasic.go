package main

import (
	"fmt"
	"runtime"
	"time"
)

// goroutines vs OS Threads
// Lighter weight (2k compare to 1M way smaller than threads)
// Go manages goroutines
// Less switching, other goroutines will run the same threads
// Safe communication

// GO's concurrency model use Actor model which is a Communicating Sequential Process

func main() {
	godur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(3)
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello", i)
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go", i)
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("2s")
	time.Sleep(dur)
}
