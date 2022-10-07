package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(num1 int) int {
	num2 := num1 % 2
	return num2
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	EvenMsg := "I have an odd number of arguments"
	OddMsg := "I have an even number of arguments"
	lengthOfArg := len(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
