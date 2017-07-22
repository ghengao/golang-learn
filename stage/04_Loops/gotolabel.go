package main

import (
	"fmt"
)

// for <pre-statements;<expression>;<post-statements> {...}
// for i:= range collection {}

func main() {
	x := 0
	for timer := 100; timer >= 0; timer-- {
		for stimer := 9; stimer >= 0; stimer-- {
			fmt.Println(timer, stimer)
			if timer%stimer == 0 {
				goto L1
			} else if timer%stimer == 3 {
				goto L2
			}
		}
	}
L1:
	fmt.Println("L1:", x)
L2:
	fmt.Println("L2:", x)

}
