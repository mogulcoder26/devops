package main

import (
	"fmt"
)

// 	"time"
// )

func main() {
	fmt.Println("Welcome to Time study of Golang!")
	// currTime := time.Now()
	// fmt.Printf("%T",currTime)
	// a := 8
	// ptr := &a
	// var ptr2 *int = &a
	// *ptr2 = *ptr2 + 2
	// fmt.Println("Value of ptr = ", ptr)
	// fmt.Println("Value of *ptr = ", *ptr)
	// fmt.Println("Value of *ptr2 = ", *ptr2)

	var fruitList [3]string
	var vegList = [3]string{"Brinjal", "Guava"}
	fmt.Println(vegList[1])
	fruitList[1] = "gon"
}
