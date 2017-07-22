package main

import (
	"fmt"
)

// NOTE: We can't change constants once it is assigned
const (
	Enforcing = 1
)

func main() {
	fmt.Println("Enforcing:", Enforcing)
	// NOTE: Compile error
	// cannot assign to Enforcing
	// cannot use 2 (type int) as type untyped number in assignment
	// Enforcing = 2
	// fmt.Println("Enforcing:",Enforcing)
}
