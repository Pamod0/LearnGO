package main

import "fmt"

func main() {
	fmt.Println("Welcome to slices in Go!")

	var mySlice []string
	fmt.Printf("Type of mySlice is %T\n", mySlice)
	fmt.Println("mySlice is: ", mySlice)

	var mySlice1 = []string{"Hello", "World", "This", "is", "Go", "Programming"}
	fmt.Println("mySlice1 is: ", mySlice1)

	mySlice1 = append(mySlice1, "Welcome", "to", "Go")
	fmt.Println("mySlice1 is: ", mySlice1)

	mySlice1 = append(mySlice1[1:5])
	fmt.Println("mySlice2 is: ", mySlice1)

	// remove an element from slice

	var courses = []string{"Python", "Java", "Go", "JavaScript", "Ruby", "C++"}

	fmt.Println("courses are: ", courses)
}