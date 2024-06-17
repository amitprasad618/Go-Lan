package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Here we learn date style in different format")

	currentTime := time.Now()
	fmt.Println("Without Formatted Time: ", currentTime)
	formattedTime := currentTime.Format("2006-01-02 3:04 PM")
	fmt.Println("Formatted Time: ", formattedTime)

	dateStr := "2024-11-29"
	parsedTime, err := time.Parse("2006-01-02", dateStr)

	if err == nil {
		fmt.Println("Parse Time: ", parsedTime)
	} else {
		fmt.Println("error parsing time", err)
	}
}
