package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time app")

	presentTime := time.Now()
	fmt.Println("Formatted time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 15:04:05 Monday"))
}