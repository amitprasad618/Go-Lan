package main

import "fmt"

func main() {
	fmt.Println("If else Switch in GO lan")

	z := 7

	if z > 10 && z < 20 {
		fmt.Println("Z is in between 10 to 20")
	} else if z > 0 && z < 9 {
		fmt.Println("Z is in between 0 to 9")
	} else {
		fmt.Println("Z is grteater 20")
	}

	day := 5

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	default:
		fmt.Println("not Monday and Not Tuesday")
	}

	month := "Jan"

	switch month {
	case "Jan", "Feb", "Mar":
		fmt.Println("1-3 month")
	case "Apr", "may", "jun", "july":
		fmt.Println("4-7month")
	default:
		fmt.Println("7 onward month")
	}

	temp := 5

	switch {
	case temp >= 25:
		fmt.Println("hot")
	case temp < 25:
		fmt.Println("cool")
	}
}
