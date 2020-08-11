package main 

import (
	"fmt"
)
// Basic array stuff
func main() {
	grades := [3]int{97, 96, 92}
	fmt.Printf("Grades: %v, %v. %v\n", grades)

	var peoples [3]string
	peoples[0] = "geohotz"
	fmt.Printf("Students: %v\n", peoples)

	cars := [3]string{"tesla", "bmw", "lambo"}
	fmt.Println(cars)

	// Go uses each thing in an array as its own variable

}