package main

import "fmt"

func modifyByRefrence(num *int) {
	*num *= 2
}

func main() {
	fmt.Println("Pointer is similar to CPP")
	var num int = 2
	var ptr *int = &num

	fmt.Println("Value of Num: ", num)
	fmt.Println("vlaue of num through Pointer: ", *ptr)

	// Nil Pointer => Pointer are initilized with nil by default if not explixity set to a point to a valid memory address.
	//				  A nil pointer doesn't point to any valid memory address

	// var ptr1 *int

	// if ptr1 == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	// modify by refrence
	value := 10
	modifyByRefrence(&value) //we are sending the address
	fmt.Println("Modified value is:", value)
}
