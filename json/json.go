package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"IsAdult"`
}

func main() {
	fmt.Println("Will Learn about Json in Go Lan")

	//Encoding Marshalling
	person := Person{Name: "Amit", Age: 28, IsAdult: true}
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error Marshalling Json", err)
		return
	}
	fmt.Println("JSON DATA: ", string(jsonData))
	// JSON DATA:  {"name":"Amit","age":28,"IsAdult":true}

	//Decoding UnMarshalling
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error in UnMarshalling", err)
		return
	}
	fmt.Println("Decode Person: ", newPerson)
	// Decode Person:  {Amit 28 true}

}
