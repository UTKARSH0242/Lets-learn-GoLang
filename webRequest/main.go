package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://utkarsh.dev"

func main()  {
	fmt.Println("Welcome to web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n ", response)

	defer response.Body.Close() // caller resposnibilty to close the connection

	databytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)

}