package main

import (
	"fmt"
)

func main() {
	var i int = 42
	fmt.Println(i)
	// OR you can do this which is so much better
	j := 22
	fmt.Println(j)
	fmt.Printf("%v, %T\n", i, i)

	// convert a string to bytes ig
	meme := "wow"
	bigMeme := []byte(meme)
	fmt.Println(bigMeme)
	// outputs array
}