package main

import "fmt"

// Pointers in Go are variables that store the memory address of another variable.
// They are used to reference and manipulate the value stored at that memory address directly.

func main() {
	var ptr *int
	fmt.Println("\nValue of ptr is: ", ptr)

	var name *string
	_name := "Hello"
	name = &_name
	fmt.Println("\nMemory address of _name: ", name)
	fmt.Println("Value of _name: ", _name)

	myNumber := 23
	var myPtr = &myNumber

	fmt.Println("\nMemory address of pointer: ", myPtr)
	fmt.Println("Value of actual pointer is: ", *myPtr)

	*myPtr = *myPtr + 2
	fmt.Println("\nNew value is: ", myNumber)
}
