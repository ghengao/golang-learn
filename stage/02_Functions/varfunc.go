package main

import "fmt"

func main() {
	bestFinish := biggest(13, 10, 13, 17, 14, 16, 20, 32)
	fmt.Println(bestFinish)
}

func biggest(finishes ...int) (best int) {
	//NOTE Passing the list of slice
	best = finishes[0]
	for _, i := range finishes {
		if i > best {
			best = i
		}
	}
	return
}
