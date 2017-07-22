package main

import "fmt"

// NOTE:to make a slice use make(<type>, <len>, <cap>)

func main() {
	myCourse := make([]int, 5, 10) //NOTE: Doesn't initialize the slice values
	fmt.Println(myCourse)
	fmt.Printf("Length is %d.\nCpacity is: %d\n", len(myCourse), cap(myCourse))
	anotherCourse := []string{"Docker", "Windows", "Linux"} //NOTE: Initialize the slice without make and assign values in one line
	fmt.Println(anotherCourse)
	fmt.Printf("Length is %d.\nCpacity is: %d\n", len(anotherCourse), cap(anotherCourse))
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(mySlice[4])
	newSlice := mySlice[2:5]
	fmt.Println(newSlice)

	//Expending a slice
	appendedSlice := append(newSlice, 1)
	//NOTE: Double the capcity when the original capacity is not enough in go.
	fmt.Printf("Length is %d.\nCpacity is: %d\n", len(appendedSlice), cap(appendedSlice))

	for _, i := range appendedSlice {
		fmt.Println("For range loop:", i)
	}
	// NOTE: How to append a slice to another slice
	myNewSlice := append(mySlice, newSlice...)
	fmt.Println(myNewSlice)

}
