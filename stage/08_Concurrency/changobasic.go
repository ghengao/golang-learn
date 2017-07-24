package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "There are the imes that try men's souls.\n"
	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}
	close(ch) // Just preventing anything input into the channel
	//for {
	//	if msg,ok := <-ch;ok {
	//		fmt.Print(msg + " ")
	//	} else {
	//		break
	//	}
	//}
	for msg := range ch {
		fmt.Print(msg + " ")
	}
	//ch<-"adding string to a close channel"
}
