package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var table []int
	tempa := 0
	tempb := 0
	temp := 0
	if n == 0 {
		z01.PrintRune(rune('0'))
	}
	for i := 0; n != 0; i++ {
		temp = n % 10
		n = n / 10
		table = append(table, 48+temp)
	}
	for i := 0; i < len(table); i++ {
		for i := 0; i < len(table)-1; i++ {
			if table[i] > table[i+1] {
				tempa = table[i]
				tempb = table[i+1]
				table[i] = tempb
				table[i+1] = tempa
			}
		}
	}
	for i := 0; i < len(table); i++ {
		z01.PrintRune(rune(table[i]))
	}
}
