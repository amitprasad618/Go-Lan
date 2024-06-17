package main

import "fmt"

func add(a, b int) (res int) {
	res = a + b
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0 {

		return 0, fmt.Errorf("invalid division") // Error message starts with lowercase
	}

	return a / b, nil
}

func main() {
	fmt.Println("Hi we are learning function")
	sum := add(2, 4)
	fmt.Println("Sum is", sum)

	div, _ := divide(4, 0)
	fmt.Println("Division is", div)
	fmt.Printf("Division %.3f", div)
}
