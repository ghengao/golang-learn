package main

import (
	"fmt"
	"math/rand"
	"time"
)

//switch <simple statement>;<expression> {
// 	case <expression> : <code>
//	case <expression> : <code>
//	default: <code>
// }

func main() {
	switch tmpNum := random(); tmpNum {
	case 0, 2, 4, 6, 8:
		fmt.Println("We got even number:", tmpNum)
		fallthrough
	case 1, 3, 5, 7, 9:
		fmt.Println("We got odd number", tmpNum)
	default:
		fmt.Println("Somthing wired happends")
	}
}

func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
