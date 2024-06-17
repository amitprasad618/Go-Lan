package main

import "fmt"

func main() {
	fmt.Println("map is use to store key value!!")

	studentGrade := make(map[string]int)

	studentGrade["Amit"] = 90
	studentGrade["John"] = 80
	studentGrade["Alis"] = 70

	//Checking if key is exist or not
	grade, exist := studentGrade["Roshan"]
	fmt.Println("Is Roshan grade exist", exist)
	fmt.Println("Roshan Grade", grade) // if not present then we get zero

	for name, grade := range studentGrade {
		fmt.Printf("%s: %d\n", name, grade)
	}
}
