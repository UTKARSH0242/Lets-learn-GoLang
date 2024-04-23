package main

import (
	"fmt"
	"sort"
)



func main()  {
	fmt.Println("Welcome to the Slices")

	var fruitList = []string{"apple","grapes", "oranges"}
	fmt.Printf("Type of fruitlist is %T\n",fruitList)


	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList [1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList [1:3])
	fmt.Println(fruitList)


	highScores := make([]int,4)


	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 123
	highScores[3] = 432

	highScores = append(highScores, 5555,66,6677,3334)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	
}