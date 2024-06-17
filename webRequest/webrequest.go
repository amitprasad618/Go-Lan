package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("will learn about web requests")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	fmt.Printf("type of response: %T\n", response)
	if err != nil {
		fmt.Println("Error getting in GET request: ", err)
		return
	}
	defer response.Body.Close()

	// Read the response body
	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error in reading response", err)
		return
	}

	fmt.Println(string(data))
}
