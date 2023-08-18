package main

import "fmt"


const LoginToken string = "aojfneg"



func main() {
	var username string = "Soubhik"
	var isLoggedIn bool = true

	fmt.Println("Name = " + username)
	fmt.Printf("Variable is of type : %T \n", username)

	fmt.Println("IsLoggedIn =", isLoggedIn)
	fmt.Printf("isLoggediIn is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255  // cant be 256.
	var smallVal2 int = 25534 // cant be 256.
	fmt.Println(smallVal)
	fmt.Println(smallVal2)
	fmt.Println(smallVal)

	var smallFloat float32 = 255.4524
	fmt.Println("smallFloat", smallFloat)

	//default vals :
	var anotherVar int

	fmt.Println(anotherVar)
	fmt.Println("AnotherVariable = ", anotherVar) //INITIALIZED TO 0
	fmt.Println("Value of AV = ", anotherVar)

	var website = "string " //inferred
	fmt.Println(website)	

	fmt.Println(LoginToken)

	numberOfUser:=3000.0090 //allowedd inside functions,not globally
	numberOfUser =3300
	fmt.Println(numberOfUser)



}
