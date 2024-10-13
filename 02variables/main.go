package main

import "fmt"

const LoginToken string = "asdasdasd" // public

func main() {
	var x string = "Hello World"
	fmt.Println(x)
	fmt.Printf("Type: %T \n", x)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type: %T \n", smallVal)

	var smallFloat float32 = 255.45555
	fmt.Println(smallFloat)
	fmt.Printf("Type: %T \n", smallFloat)

	// default values and some aliases

	var intVariavle int
	fmt.Println(intVariavle)
	fmt.Printf("Type: %T \n", intVariavle)

	// implicit type

	var name = "Pamod"
	fmt.Println(name)

	// no var style

	numberOfUsers := 300
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Type: %T \n", LoginToken)
}
