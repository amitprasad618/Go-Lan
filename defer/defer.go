package main

import "fmt"

func main() {
	fmt.Println("defer will excuate in last when main func excuate")
	defer fmt.Println("Middle of Program")
	fmt.Println("end of Prog")
	defer fmt.Println("end->next of Prog")

	// defer will Work as LIFO pattern

	// defer we are using to close file after using.
	// kahi v defer likho last me excuate hoga LIFO format me
}
