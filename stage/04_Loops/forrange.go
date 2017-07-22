package main

import "fmt"

func main() {
	courseInProg := []string{"Docker", "Linux", "Windows"}
	courseCompleted := []string{"Docker", "Windows"}
	for _, course := range courseInProg {
		//fmt.Println(index, course)
		for _, course_completed := range courseCompleted {
			if course == course_completed {
				fmt.Println("\nHey, we found a complete course", course)
			}
		}
	}
}
