package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Will convert data from one type to other type")

	var integerNum int = 45
	var floatNum float64 = float64(integerNum)
	fmt.Printf("integerNum to floatNum: %.3f\n", floatNum)
	fmt.Println("integerNum to floatNum: ", floatNum)

	var num int = 200
	str := strconv.Itoa(num)
	fmt.Printf("str is %s\n", str)
	fmt.Printf("Type of str is %T\n", str)

	// str1 := "3.14"
	// num, err := strconv.ParseInt(str1, 10, int64)
	// if err != nil {
	// 	fmt.Println(num)
	// } else {
	// 	fmt.Println("Error:", err)
	// }

}
