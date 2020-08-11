package main

import (
	"fmt"
)

func main() {
	nerd := []int{1, 2, 3, 4, 7, 9}

	fmt.Println(len(nerd)) // length
	fmt.Println(cap(nerd)) // capacity

	fmt.Println(nerd[1:3])

	other := make([]int, 3, 100) // array of ints with 3 elements and a capacity of 100	
	fmt.Println(other)

	for _, num := range nerd {
		println(num)
	}

}