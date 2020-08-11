package main 

import (
	"fmt"
)

// Basically this is a predefined map structure
type Student struct {
	id int
	studentName string
	grades []int
}

func main() {
	studentOne := Student{ // Here I'm creating a map of studentOne using the student structure that we defined already
		id: 223,
		studentName: "John Hammond",
		grades: []int{97, 91, 93, 89, 90},
	}
	fmt.Println(studentOne) // Prints our the studentOne map
	fmt.Println(studentOne.id) // Prints out the value of id in the studentOne map
}