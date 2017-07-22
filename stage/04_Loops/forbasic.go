package main

import (
	"fmt"
	"time"
)

// for <pre-statements;<expression>;<post-statements> {...}
// for i:= range collection {}

func main() {

	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom!")
			break
		}
		fmt.Println(timer)
		time.Sleep(1 * time.Second)
	}

}
