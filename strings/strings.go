package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("there is lots of inbuild func for manage string in GO Lan")
	str := "Apple,Mango,Banana"

	parts := strings.Split(str, ",")
	fmt.Println(parts)

	str1 := "one two one two two two three four two"

	count := strings.Count(str1, "two")
	fmt.Println("the count of two in str1 is: ", count)

	str2 := "Hello"
	str3 := "world"

	res := strings.Join([]string{str2, str3}, " ")
	fmt.Println("After concinating str2 and str3: ", res)
}
