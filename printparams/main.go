package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	var str string
	for j := 1; j < len(arguments); j++ {
		str = string(arguments[j])
		argu := []rune(str)
		for i := 0; i < len(argu); i++ {
			z01.PrintRune(argu[i])
		}
		z01.PrintRune('\n')
	}
}
