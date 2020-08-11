package main 

import (
	"fmt"
)

func main() {
	n := 1 == 1 // Prints true or false based on equivilancy test
	m := 2 == 3
	fmt.Println(n)
	fmt.Println(m)

	var meme bool // Will initialize as 0 by default
	fmt.Printf("%v, %T\n", meme, meme)

	// Types of integers
	// int8, int16, int32, int64

	// aaa
	var lol uint16 = 42
	fmt.Println(lol)
}