package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Go!")

	var myArray [5]string
	myArray[0] = "Hello"
	myArray[1] = "World"
	myArray[2] = "This"
	myArray[3] = "is"
	myArray[4] = "Go"

	fmt.Println("myArray is: ", myArray)
	fmt.Println("myArray is: ", len(myArray))

	var vegetables = [3]string{"potato", "tomato", "onion"}
	fmt.Println("Vegetables are: ", vegetables)
}