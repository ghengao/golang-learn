package main

import (
	"fmt"
	"reflect"
)

// Package Variables Declaration
//var (
//	name, course string
//	module       float64
//)

// Package Variables Declaration and Initialization in one line
//var (
//	name, course, module = "Heng GAO","Docker Deep Dive",3.2
//)

// Best/Clean way to declare and initialize the variables on each line
var (
	name   = "Heng GAO"         //name
	course = "Docker Deep Dive" //course
	module = 3.2                //module
)

func main() {
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is set to", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is set to", module, "and is of type", reflect.TypeOf(module))
}
