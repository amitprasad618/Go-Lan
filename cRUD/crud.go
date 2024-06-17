package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   //No JSON TAG specified
}

func PerformGETRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch data: ", response.Status)
		return
	}

	//Reading Json data from GET Request
	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("error in reading response", err)
	// 	return
	// }
	// fmt.Println(string(data))

	//Decoding GET Request
	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("error decoding response", err)
		return
	}

	fmt.Println("Todo Id: ", todo.ID)
	fmt.Println("Todo Title: ", todo.Title)
	fmt.Println("Todo Completed: ", todo.Completed)
}

func PerformPOSTRequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos"

	todo := Todo{
		UserId:    1,
		ID:        201,
		Title:     "Learn GO",
		Completed: false,
	}

	//Convert the todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error in Json Marshalling", err)
		return
	}

	//Convert the JSON byte slice to a string
	jsonStr := string(jsonData)

	//Create an io.Reader from the string
	jsonReader := strings.NewReader(jsonStr)

	//Send POST Request
	resp, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending Post req", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)
}

func PerformUPDATERequest() {
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	todo := Todo{
		UserId:    1,
		ID:        1,
		Title:     "Update Todo",
		Completed: true,
	}

	//Convert the todo struct to Json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("ERROR in Marshalling JSON", err)
		return
	}

	//Create PUT req
	req, err := http.NewRequest(http.MethodPut, myUrl, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error Creating req", err)
		return
	}
	req.Header.Set("content-type", "application/json")

	//send request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response Status", res.Status)

}

func PerformDELETERequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//Create Delete Request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error Creating Request", err)
		return
	}

	//send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending Request", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status", resp.Status)
}

func main() {
	fmt.Println("Will about CRUD Opration GET, POST, PUT, DELETE")
	fmt.Println("--------------------------------------------")
	fmt.Println("GET Request")
	PerformGETRequest()
	fmt.Println("--------------------------------------------")
	fmt.Println("POST Request")
	PerformPOSTRequest()
	fmt.Println("--------------------------------------------")
	fmt.Println("PUT Request")
	PerformUPDATERequest()
	fmt.Println("--------------------------------------------")
	fmt.Println("DELETE Request")
	PerformDELETERequest()
}
