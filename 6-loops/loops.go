package main

import "fmt"

func main() {
	dc := []string{"batman", "justice league", "wonder women"}
	fmt.Println(dc)

	// this might be confusing, but go uses for syntax for while loop also
	i := 0
	for i < 3 {
		fmt.Printf("array value for index %v is %v \n", i, dc[i])
		i++
	}
}
