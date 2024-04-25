package main

import "fmt"
func main()  {
	fmt.Println("Defer in golang")
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("four")
	mydefer()
}

func mydefer(){
	for i := 0; i < 5; i++ {
		// defer fmt.Print(i +" ")
		defer fmt.Print(fmt.Sprint(i) + " ")
	}
}