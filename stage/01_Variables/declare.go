package main

import (
	"fmt"
	"reflect"
)

// NOTE:Package Variables Declaration
//var (
//	name, course string
//	module       float64
//)

// NOTE:Package Variables Declaration and Initialization in one line
//var (
//	name, course, module = "Heng GAO","Docker Deep Dive",3.2
//)

// NOTE:Best/Clean way to declare and initialize the variables on each line
var (
	name   = "Heng GAO"         //name
	course = "Docker Deep Dive" //course
	module = 3.2                //module
)

// NOTE:Go does not allow := in in the package variable declaration
//var (
//	name   := "Heng GAO"         //name
//	course := "Docker Deep Dive" //course
//	module := 3.2                //module
//)
//.\declare.go: syntax error: unexpected :=

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is set to", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
	//a := 10
	//b := 3
	//c := a + b
	//fmt.Println(c)
	//13

	//a:=10.00000
	//b:=3
	//c:=a+b
	//fmt.Println(c)
	// NOTE:.\integer.go:14: invalid operation: a + b (mismatched types float64 and int)
	// NOTE:However it allow type conversion
	// := only available within the function when initialize the variable.
	// NOTE::If you declare the variable within the function but not use it go will raise compile exception.
	// NOTE::However golang allows you to declare the variable within the package without using it.
	a := 10.000000
	b := 3
	fmt.Println("\nA is type", reflect.TypeOf(a), "and B is of type", reflect.TypeOf(b))
	c := int(a) + b
	// NOTE: A does not change type, but c follows the type of the operation.
	fmt.Println("\nC is type", reflect.TypeOf(c), "and A is of type", reflect.TypeOf(a))
}
