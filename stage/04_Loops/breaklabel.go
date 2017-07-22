package main

import (
	"fmt"
)

// for <pre-statements;<expression>;<post-statements> {...}
// for i:= range collection {}

func main() {
L1:
	for timer := 100; timer >= 0; timer-- {
		for stimer := 9; stimer >= 0; stimer-- {
			fmt.Println(timer, stimer)
			if timer%stimer == 0 {
				break L1
			}
		}
	}

}
