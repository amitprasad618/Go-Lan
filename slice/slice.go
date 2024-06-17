package main

import "fmt"

func main() {
	fmt.Println("Slice is nothing but Vector in CPP")

	numbers := []int{1, 2, 3, 3, 4, 4}
	fmt.Println("number slice is", numbers)
	fmt.Println("Length of numbers slice is", len(numbers))
	fmt.Println("Capacity of numbers slice is", cap(numbers))

	//USE of Make function to create a slice with a specific length and capacity

	score := make([]int, 3, 5) //type length capacity
	score = append(score, 22)
	score = append(score, 232)
	score = append(score, 224)
	fmt.Println("score slice is", score)
	fmt.Println("Length of score slice is", len(score))
	fmt.Println("capacity of score slice is", cap(score))
}
