package main

import "fmt"

//switch <simple statement>;<expression> {
// 	case <expression> : <code>
//	case <expression> : <code>
//	default: <code>
// }

func main() {
	switch "docker" {
	case "linux":
		fmt.Println("\nHere 1")
	case "docker":
		fmt.Println("\nHere 2")
	case "windows":
		fmt.Println("\nHere 3")
	default:
		fmt.Println("\nSorry, no match")
	}
}
