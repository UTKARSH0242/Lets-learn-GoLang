package main

import "fmt"

const LoginToken string = "naicaddv" // public

func main() {
	// fmt.Println("this is my variable")
	var userName string = "utkarsh"
	fmt.Println(userName)
	fmt.Printf("variable is of type: %T \n", userName)

	var isBool bool = true
	fmt.Println(isBool)
	fmt.Printf("variable is of type: %T \n", isBool)

	numberOfUser := 98754.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
