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

	// slice, uses array under the hood, dynamic length, length is infered, compiler decides the length based on values
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

	/** Slicing */
	arr6 := [8]int{11, 15, 19, 23, 29, 34, 40, 55}
	// slice array, 1st element to 3rd element including
	newslicedarr := arr6[1:4]
	fmt.Println("Length:", len(arr6))
	fmt.Println("Array :", arr6)
	fmt.Println("Sliced:", newslicedarr)

	// slice array, 4th and after 4th element
	newslicedarr2 := arr6[4:]
	fmt.Println("Length:", len(arr6))
	fmt.Println("Array :", arr6)
	fmt.Println("Sliced:", newslicedarr2)

	// append, it makes a copy of new array, so we have to assign it
	arr7 := []int{5, 6, 7}
	arr8 := append(arr7, 8)
	fmt.Println(arr8)

}
