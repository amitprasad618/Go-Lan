package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name: ")
	// var name string
	// fmt.Scan(&name)  //we are not able to take while we give FULL NAME i.e. Amit Prasad

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')

	fmt.Println("My name is:", name)
}
