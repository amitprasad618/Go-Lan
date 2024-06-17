package main

import "fmt"

func main() {
	fmt.Println("Only one Loop Present in Go lan i.e. For")

	for i := 0; i < 3; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println(" ")
	//use as while loop
	counter := 10

	for {
		fmt.Print(" ", counter)
		counter++
		if counter == 15 {
			break
		}
	}

	fmt.Println(" ")
	fmt.Println("-----------------------------------------------------------------------------")
	//Range Keyward use to print like key value ==> index in loop, key in Map
	num := []int{1, 2, 3, 4, 5, 6, 7}
	for index, val := range num {
		fmt.Printf("Index is %d and value is %d\n", index, val)
	}
}
