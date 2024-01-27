package main

import "fmt"

func main() {
	dc := []string{"batman", "justice league", "wonder women"}
	fmt.Println(dc)

	// while loop - this might be confusing, but go uses for syntax for while loop also
	i := 0
	for i < 3 {
		fmt.Printf("array value for index %v is %v \n", i, dc[i])
		i++
	}

	// normal for loop, same as other languages
	for i := 0; i < len(dc); i++ {
		fmt.Printf("- The array value for index %v is %v \n", i, dc[i])
	}

	// for loop for array, using index,value pair, same as python, foreach in php
	for index, value := range dc {
		fmt.Printf("--- The array value for index %v is %v \n", index, value)
	}

	// if you dont want any unused variable like index or value, use underscore
	for _, value := range dc {
		fmt.Printf("------The array value is %v \n", value)
	}
}
