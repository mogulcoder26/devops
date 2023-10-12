package main

import "fmt"

func main() {
	// fmt.Println("Structs in Golang!")

	soubhik := User{"Soubhik", 19, false, 23}
	fmt.Println(soubhik)
	// char := soubhik.isGay
	// fmt.Println(char)
	girls := []string{"Gal", "Dua", "Disha", "Deepika"}
	for i, name := range girls {
		fmt.Printf("%d -> %s\n", i, name)
	}

}

type User struct {
	name   string
	age    int
	isGay  bool
	oneAge int
}
