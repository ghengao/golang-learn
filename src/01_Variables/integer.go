package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := 10
	//b := 3
	//c := a + b
	//fmt.Println(c)
	//13

	//a:=10.00000
	//b:=3
	//c:=a+b
	//fmt.Println(c)
	//.\integer.go:14: invalid operation: a + b (mismatched types float64 and int)
	//However it allow type conversion

	a := 10.000000
	b := 3
	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type,", reflect.TypeOf(b))

}
