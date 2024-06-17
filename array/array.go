package main

import "fmt"

func main() {
	fmt.Println("Hi we are learing Arrays")
	num := [3]int{1, 2, 36}
	fmt.Println("Array Num ", num)

	var arr [5]string //index 5 out of bounds
	arr[0] = "Apple"
	arr[1] = "Mango"
	arr[4] = "banana"

	fmt.Println("Array Arr ", arr)

	var score [5]int
	fmt.Println("Score array", score)
}
