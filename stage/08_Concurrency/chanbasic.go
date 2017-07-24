package main

import "fmt"

//

func main() {
	// Deadlock example:
	//ch := make(chan string)
	//fmt.Println(<-ch)
	// Another Deadlock, because channel does not have the capacity.
	//ch := make(chan string)
	//ch <- "Hello"
	//fmt.Println(<-ch)
	//simple use example of buffered channel.
	ch := make(chan string, 1)
	ch <- "Hello"
	fmt.Println(<-ch)
}
