package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	Street string
	City   string
}

type Contact struct {
	Email string
	Phone int
}

type Employee struct {
	Person
	Address
	Contact
	Position string
}

func main() {
	fmt.Println("Struct is similar to class in CPP")
	var per1 Person
	per1.FirstName = "Amit"
	per1.LastName = "Prasad"
	per1.Age = 28
	fmt.Println("Per1: ", per1)

	per2 := Person{
		FirstName: "Roshan",
		LastName:  "Gupta",
		Age:       29,
	}
	fmt.Println("Per2: ", per2)

	//Struct Embedding within other structs, allowing for a form composition
	emp := Employee{
		Person: Person{
			FirstName: "Jyoti",
			LastName:  "Gupta",
			Age:       26,
		},
		Address: Address{
			Street: "2323",
			City:   "Deoria",
		},
		Contact: Contact{
			Email: "hjyotiQ@jsdjsa.cim",
			Phone: 878721878137821,
		},
		Position: "Engineer",
	}

	fmt.Println("Employee1: ", emp)
}
