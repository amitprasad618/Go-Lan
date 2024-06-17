package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("will learn about file creating, Reading and deleting")

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error in file creating", err)
		return
	}
	defer file.Close()

	initilaContent := "HI this is Amit Prasad Learining Go Language"

	_, err = io.WriteString(file, initilaContent+"\n")
	if err != nil {
		fmt.Println("Error in writing the file: ", err)
		return
	}

	fmt.Println("File Created with initial Content")

	//Open the file
	file1, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error in opening file", err)
		return
	}
	defer file1.Close()

	//create a buffer to read the file content
	buffer := make([]byte, 1024)

	for {
		n, err := file1.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error in Reading file: ", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}

}
