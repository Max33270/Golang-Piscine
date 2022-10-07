package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	arr := make([]byte, 14)

	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) == 2 {
		file.Read(arr)
		fmt.Println(string(arr))
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	}
}
