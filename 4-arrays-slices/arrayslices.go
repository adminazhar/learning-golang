package main

import "fmt"

func main() {
	fmt.Println("hello")

	// full syntax with defined length
	var arr1 [3]int = [3]int{10, 12, 14}
	fmt.Println(arr1)

	// full syntax with defined length
	var arr2 [2]string = [2]string{"hello", "world"}
	fmt.Println(arr2)
}
