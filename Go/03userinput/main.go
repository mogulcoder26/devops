package main

import (
	"fmt"
	"time"
)

func main() {
	// 	welcome := "Welcome to User Input "
	// 	fmt.Println(welcome)

	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Println("Enter the Rating for our Pizza : ")

	// 	//comma ok syntax
	// //  try  ,  catch
	// 	input, err := reader.ReadString('\n')
	// 	fmt.Println("Thanks for Rating ", input)
	// 	fmt.Println("Err : ", err)
	// 	fmt.Printf("Err  %T\n: ", err)
	// 	fmt.Printf("Type of Rating  = %T \n", input)

	/*------------------------------------------------*/

	// fmt.Println("Welcome to the Pizza App")
	// fmt.Println("Please rate our pizza between 1 and 5")

	// reader := bufio.NewReader(os.Stdin)

	// input, _ := reader.ReadString('\n')
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) ;

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Added 1 to your rating : ", numRating+1)
	// }

	/*------------------------------------------------------*/

	fmt.Println("Welcome to Time study of GO")
	presentTime := time.Now()
	fmt.Println("presentTime : ", presentTime)
    fmt.Println(presentTime.Format("01-02-2006 Monday January 15:04:05"))

	// createdDate := time.Date(202)
}
