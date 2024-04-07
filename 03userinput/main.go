package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var welcome string = "I am going to provide userinput in this"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the go language rating: ")

	// comma ok // error ok

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)
	fmt.Printf("Type of rating is %T ", input)
}