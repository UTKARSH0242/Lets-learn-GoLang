package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array in go lang")

	var fruitList [4] string

	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[3] = "Banana"

	fmt.Println("Fruit list is : " , fruitList)
	fmt.Println("Fruit list is : ", len(fruitList))


	var vegList = [3]string{"patato","beans","mushroom"}
	fmt.Println("vegy list is :", len(vegList))
	
}