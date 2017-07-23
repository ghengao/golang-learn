package main

import "fmt"

// struct is customized types when builtin types are just not efficient.
// fields in a struct can be different type
// the core type for OO programming

func main() {
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}
	// Declare a struct
	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)
	// Initialize a struct
	DockerDeepDive := courseMeta{
		"Heng GAO",
		"Intermediate",
		64.4,
	}
	fmt.Println(DockerDeepDive)
	fmt.Println("Author:", DockerDeepDive.Author)
	DockerDeepDive.Rating = 1
	fmt.Println("Rating:", DockerDeepDive.Rating)

}
