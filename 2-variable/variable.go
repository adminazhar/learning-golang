package main

import "fmt"

func main() {
	/** Strings */
	var name1 string = "azhar" // full syntax fpr initializing variable
	fmt.Println(name1)

	var name2 = "vscode" // datatype is inferred from value
	fmt.Println(name2)

	name3 := "search" // shorthand, datatype inferred
	fmt.Println(name3)

	/** Integars */
	var number1 int8 = 127 // 8bit integar, range: -128 through 127
	fmt.Println(number1)

	var number2 = 254 // datatype is inferred as int
	fmt.Println(number2)

	number3 := -24324234 // inferred datatype as int
	fmt.Println(number3)

	var number4 uint16 = 234 // non-negative numbers, 0 through 65535
	fmt.Println(number4)

	/** Float */
	var decimalnum1 float32 = 2.24
	fmt.Println(decimalnum1)

	var decimalnum2 = 2.44 // datatyps is inferred, default float as float64
	fmt.Println(decimalnum2)

}
