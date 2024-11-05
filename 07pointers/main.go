package main

import "fmt"

// store memory the address of a value using pointers
// & operator is used to get the memory address of a variable
// * operator is used to get the value stored at that memory address

func incrementByValue(x int) {
	x++ // This modification is local to the function
}

func incrementByPointer(x *int) {
	*x++ // This modifies the original value outside the function
}

func main() {
	var year int = 2021
	var ptrYear *int = &year
	fmt.Println("Value of year is: ", year)
	fmt.Println("Memory address of year is: ", ptrYear)

	var ptr *int
	fmt.Println("\nValue of ptr is: ", ptr)

	myNumber := 23
	var myPtr = &myNumber

	fmt.Println("\nMemory address of pointer: ", myPtr)
	fmt.Println("Value of actual pointer is: ", *myPtr)

	*myPtr = *myPtr + 2
	fmt.Println("\nNew value is: ", myNumber)

	// Passing pointers to functions

	a := 10
	incrementByValue(a)
	fmt.Println(a) // Output: 10 (no change)

	incrementByPointer(&a)
	fmt.Println(a) // Output: 11 (a was modified)
}
