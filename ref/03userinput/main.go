package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name : ")
	input, _ := reader.ReadString('\n')

	fmt.Println("Your Name is = ", input)
}
