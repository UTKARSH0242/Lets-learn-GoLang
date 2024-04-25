package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main()  {
	fmt.Println("Welcome to the file system in golang")
	content := "This needs to go in a file - file.txt"
	file, err := os.Create("./file.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./file.txt")
}

func readFile(filename string){
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is: \n", string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}