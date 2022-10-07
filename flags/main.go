package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	if len(arg) < 2 || arg[1] == "-h" || arg[1] == "--help" {
		s := "--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n"
		for _, w := range s {
			z01.PrintRune(w)
		}
	}
	if len(arg) > 2 && IsInsert(arg[1]) && !(IsOrder((arg[2]))) {
		str := ""
		a := []rune(arg[1])
		b := 0
		for i := 0; i < len(a); i++ {
			if a[i] == 61 {
				b = i
			}
		}
		for i := b + 1; i < len(arg[1]); i++ {
			str = str + string(a[i])
		}
		str = arg[2] + str
		for _, w := range str {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
	if len(arg) > 2 && IsOrder(arg[1]) {
		a := []rune(arg[2])
		for i := 0; i < len(a); i++ {
			for i := 0; i < len(a)-1; i++ {
				if a[i] > a[i+1] {
					a[i], a[i+1] = a[i+1], a[i]
				}
			}
		}
		for i := 0; i < len(a); i++ {
			z01.PrintRune(rune(a[i]))
		}
		z01.PrintRune('\n')
	}
	if len(arg) > 2 && IsOrder(arg[2]) && IsInsert(arg[1]) {
		str := ""
		a := []rune(arg[1])
		y := 0
		for i := 0; i < len(a); i++ {
			if a[i] == 61 {
				y = i
			}
		}
		for i := y + 1; i < len(arg[1]); i++ {
			str = str + string(a[i])
		}
		str = arg[3] + str
		b := []rune(str)
		for i := 0; i < len(str); i++ {
			for i := 0; i < len(str)-1; i++ {
				if b[i] > b[i+1] {
					b[i], b[i+1] = b[i+1], b[i]
				}
			}
		}
		for i := 0; i < len(b); i++ {
			z01.PrintRune(rune(b[i]))
		}
		z01.PrintRune('\n')
	}
}

func IsInsert(s string) bool {
	a := []rune(s)
	if a[0] == '-' && a[1] == 'i' {
		return true
	} else if a[0] == '-' && a[1] == '-' && a[2] == 'i' && a[3] == 'n' && a[4] == 's' && a[5] == 'e' && a[6] == 'r' && a[7] == 't' {
		return true
	} else {
		return false
	}
}

func IsOrder(s string) bool {
	if s == "--order" || s == "-o" {
		return true
	} else {
		return false
	}
}
