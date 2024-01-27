package main

import (
	"fmt"
	"strings"
)

func main() {
	quote := "hello world!"
	// Checks if string subtring is in the string
	fmt.Println(strings.Contains(quote, "world"))

	// count substr occurence
	fmt.Println(strings.Count(quote, "l"))

	// check index positionn
	fmt.Println(strings.Index(quote, "o"))
}
