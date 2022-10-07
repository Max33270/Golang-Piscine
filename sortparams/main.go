package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	runes := os.Args[1:]
	BubbleSort(runes)
	for _, v := range runes {
		for _, t := range v {
			z01.PrintRune(t)
		}
		z01.PrintRune('\n')
	}
}

func BubbleSort(array []string) []string {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
