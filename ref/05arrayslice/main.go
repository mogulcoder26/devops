package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello to Arrays and Slices!")

	// arr := [5]int{1, 2, 45, 23}
	// fmt.Println(arr)

	// fruitList := []string {"Apple","Tomato","Peach"};
	// fmt.Println(fruitList);
	// fruitList = append(fruitList, "Lund")
	// fmt.Println(fruitList);
	// fruitList = fruitList[1:2]

	arr := []string{"one", "two", "three", "four", "five"}
	arr = append(arr[:3], arr[4:]...)
	fmt.Println(arr)

	hoes := make(map[string]int, 5)
	hoes["Sid"] = 1000
	hoes["Soubhik"] = 0

	fmt.Println(hoes)

}
