package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	_, err := os.Open("c:\\temp\\test1.txt")
	if err != nil {
		fmt.Println("Error returned was:", err, "Type of err:", reflect.TypeOf(err))
	}
}
