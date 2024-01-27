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

	// replace, returns a copy, find and replace
	fmt.Println(strings.Replace(quote, "hello", "hi", 1))

	// replaceall, same as replace but dont need how many times
	fmt.Println(strings.ReplaceAll(quote, "hello", "yo"))

}
