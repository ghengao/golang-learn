package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()
	logCh := make(chan string, 50)
	go func() {
		for {
			msg, ok := <-logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + "-" + msg)
				f.Close()
			} else {
				break
			}
		}
	}()

	mutex := new(sync.Mutex)

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			mutex.Lock()
			go func(i, j int) { // Try to remove the passing parameters and let the goroutines to share the variable i,j from the main routine.
				msg := fmt.Sprintf("%d + %d =%d\n", i, j, i+j)
				logCh <- msg
				fmt.Printf("%d + %d =%d\n", i, j, i+j)
				mutex.Unlock()
			}(i, j)
		}
	}
	fmt.Scanln()
}
