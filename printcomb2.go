package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	i := 0
	j := 0
	k := 0
	l := 0
	for a := 0; a < 10000; a++ {
		if (i*10)+j < (k*10)+l {
			z01.PrintRune(rune(i + 48))
			z01.PrintRune(rune(j + 48))
			z01.PrintRune(' ')
			z01.PrintRune(rune(k + 48))
			z01.PrintRune(rune(l + 48))
			if i == 9 && j == 8 && k == 9 && l == 9 {
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		l += 1
		if l == 10 {
			l = 0
			k += 1
		}
		if k == 10 {
			k = 0
			j += 1
		}
		if j == 10 {
			j = 0
			i += 1
		}
	}
	z01.PrintRune('\n')
}
