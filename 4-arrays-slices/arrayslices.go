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

	// infered
	var arr3 = [4]int{45, 47, 49, 99}
	fmt.Println(arr3)

	// shorthand
	arr4 := [2]string{"batman", "robin"}
	fmt.Println(arr4)

	// length is infered, compiler decides the length based on values
	arr5 := []int{18, 19, 20}
	fmt.Println(arr5)

	commonarr := []string{"batman", "dark knight", "dark knight rises"}
	// accessing elements
	fmt.Println(commonarr[0])
	fmt.Println(commonarr[1])
	// updating element
	commonarr[1] = "the dark knight"
	fmt.Println(commonarr)
	// count length
	println(len(commonarr))

}
