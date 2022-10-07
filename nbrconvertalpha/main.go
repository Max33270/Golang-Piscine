package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		if Atoi(args[i]) > 26 {
			z01.PrintRune(32)
		}
		if args[1] == "--upper" && Atoi(args[i]) >= 1 && Atoi(args[i]) <= 26 {
			z01.PrintRune(rune(Atoi(args[i]) + 64))
		} else if Atoi(args[i]) >= 1 && Atoi(args[i]) <= 26 {
			z01.PrintRune(rune(Atoi(args[i]) + 96))
		} else if Atoi(args[i]) == 0 && args[i] != "--upper" {
			z01.PrintRune(32)
		}

	}
	if len(os.Args) > 1 {
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	// define 2 variables
	var resultat int = 0
	var compte int = 1
	// loop to check all elements in s backwards, starting from the end
	for i := len(s) - 1; i >= 0; i-- {
		// if the element is a + sign we return s
		if s[i] == 43 && i == 0 {
			return resultat
			// if the element is a - (minus) sign we return s
		}
		if s[i] == 45 && i == 0 {
			resultat *= -1
			return resultat
			// third check : if it's a number between 0 and 9 we return s
		}
		if s[i] < 48 || s[i] > 57 {
			return 0
		}
		resultat += int(s[i]-48) * compte
		compte = 10 * compte
	}
	return resultat
}
