package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := make([]byte, 200)
	ptr := unsafe.Pointer(&s[0])
	fmt.Println(ptr)
}
