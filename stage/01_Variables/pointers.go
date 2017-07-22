package main

import (
	"fmt"
	"reflect"
)

// NOTE::Go pass arguments by value
//

func main() {
	name := "Heng GAO"
	course := "Docker Deep Dive"
	module := 3.2
	//NOTE: & gives the reference pointer with the memory address of the variable
	//NOTE: * De-reference a pointer
	fmt.Println("Name is set to", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Memory address of *module variable is", &module)
	fmt.Println("Type of the Memory address of *module is", reflect.TypeOf(&module))
	fmt.Println("Dereference the module", *&module)

	// NOTE: Go default make a copy of the arguments instead of pointers unless we give the pointer to it.
	fmt.Println("\nHi, You are watching:", course)
	changeCourse(course)
	fmt.Println("\nTrying to change your course to:", course)
	changeCourseByReference(&course)
	fmt.Println("\nTrying to change your course to:", course)
}
func changeCourse(course string) {
	course = "See if the argument changes"
}

func changeCourseByReference(course *string) {
	*course = "See if the argument changes"
}
