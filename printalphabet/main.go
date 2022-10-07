package main

import "github.com/01-edu/z01"

func main() {
	alphabet := "abcdefghijklmnopqrstuvwxyz\n"
	for i := 0; i < len(alphabet); i++ {
		z01.PrintRune(rune(alphabet[i])); 
	}
}
