package main

import "fmt"

func main() {
	fmt.Println("Variables in Go Lan")

	var age int = 24
	var height float64 = 5.8
	var name string = "Amit"

	fmt.Println("name:", name, " height:", height, " age:", age)

	score := 45

	fmt.Printf("type of score is %T\n", score)
	fmt.Printf("type of name is %T\n", name)

	//FORMAT Specifier
	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name is %s\n", name)
	fmt.Printf("Height is %.3f\n", height)

}
