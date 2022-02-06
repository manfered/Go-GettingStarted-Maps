package main

import "fmt"

func main() {
	// defining a map data type with string keys and integer values
	mapValue := map[string]int{"Fred": 1, "Manfered": 2}
	fmt.Println(mapValue)
	fmt.Println(mapValue["Fred"])

	// update an item's value by key
	mapValue["Fred"] = 5
	fmt.Println(mapValue["Fred"])

	// delete an item from a map by key
	delete(mapValue, "Fred")

	fmt.Println(mapValue)
}
