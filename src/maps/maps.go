package main

import (
	"fmt"
)

func main() {
	// Basically dictionaries in python
	// map[string]int
	// keyword map + key in [] + value

	classSize := map[string]int{ // key string and value of ints
		"classOne": 12,
		"classTwo": 32,
		"classThree": 22,
	}

	fmt.Println(classSize) // Print whole map

	classSize["classFour"] = 28 // Adds a new key with a new value
	fmt.Println(classSize)

	fmt.Println(classSize["classTwo"]) // Print value for key classTwo
}