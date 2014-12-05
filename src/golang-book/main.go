package main

import "fmt"

func main() {

	// printing strings
	fmt.Println("String Printing Practice\n")
	fmt.Println("1 + 2 =", 1+2)
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!"[1])
	fmt.Println("Hello, " + "World!")

	// true / false practice
	fmt.Println("\nTrue / False Practice \n")
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(!true || !false)

	// creating new variables
	fmt.Println("\nCreating new variables \n")
	x := "Hello, there!"
	y := "These variables will be concatenated into one string and the types are inferred."
	fmt.Println(x + " " + y)

	const x1 string = "Hello, this is a constant."
}
