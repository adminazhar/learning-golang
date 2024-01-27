package main

import "fmt"

func main() {
	// Print with new line
	fmt.Println("Hello world!")

	// Print without new line, \n adds new line
	fmt.Print("Hello world again! \n")

	name1 := "batman"
	age1 := 35

	// Printing directly
	fmt.Println("my name is", name1, "and age is", age1)

	// formatting with strings - formt specifier %v to replace variable
	fmt.Printf("-my name is %v and age is %v \n", name1, age1)

	// formatting - same as C lang, %s string, %d integar
	fmt.Printf("--my name is %s and age is %d", name1, age1)

}
