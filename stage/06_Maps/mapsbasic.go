package _6_Maps

import "fmt"

// creating a map using map[keyType]valueType, keyType must be comparable type(== != can be used), all keys must be unique.
// Maps are reference types, cheap to pass but unsafe for concurrency
// Specify size for large maps
// make(map[<keyType>]<valueType>,size)

func main() {
	leagueTitle := make(map[string]int)
	leagueTitle["Sutherland"] = 6
	leagueTitle["Newcastle"] = 4
	fmt.Println(leagueTitle)

	// Initialize the maps
	recentHead2Head := map[string]int{
		"Sutherland": 5,
		"Newcastle":  4,
	}
	fmt.Println(recentHead2Head)

	// Iterating a map, this suggest, go randomly store the map, every time you iterating them will be different
	testMap := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
		"D": 1,
		"E": 1,
		"F": 1,
		"G": 1,
	}
	for key, value := range testMap {
		fmt.Println(key, value)
	}
	fmt.Println("---------------")
	for key, value := range testMap {
		fmt.Println(key, value)
	}
	fmt.Println("---------------")
	for key, value := range testMap {
		fmt.Println(key, value)
	}
	fmt.Println("################")

	//How to delete a item.
	delete(testMap, "A")
	fmt.Println(testMap)

	//What happens if we add more value to a sized map
	myTestMap := make(map[string]int, 1)
	myTestMap["Test1"] = 1
	fmt.Println(testMap)
}
