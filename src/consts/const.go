package main 

import (
	"fmt"
)

// Easy way to have multiple constants
const (
	r = 3
	t = 22
	g = 3
)


func main() {
	// constant naming same as normal vars
	const coolConstant int = 3
	fmt.Printf("%v, %T\n", coolConstant, coolConstant)


}