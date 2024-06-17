package myUtil

import "fmt"

var PublicVariable int = 89  //(Exported)
var privateVariable int = 98 //(UnExported)

func PrintMessage(msg string) {
	fmt.Println(msg)
	fmt.Println("PublicVarible: ", PublicVariable)
	fmt.Println("privateVariable: ", privateVariable)
}
