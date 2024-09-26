package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23

	var myPtr = &myNumber

	fmt.Println("Value of actual pointer is: ", myPtr)
	fmt.Println("Value of actual pointer is: ", *myPtr)

	*myPtr = *myPtr + 2
	fmt.Println("New value is: ", myNumber)

}