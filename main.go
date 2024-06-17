package main

import (
	"fmt"
	myUtil "myLearning/muUtil"
)

func main() {
	fmt.Println("Hi this is Amit Prasad Learning Go Lan")
	myUtil.PrintMessage("Hi Go Lan!!")
	fmt.Println("in main access PublicVariable: ", myUtil.PublicVariable)
	//we cann't access private variable because it is in starting with small letter...!!
}
