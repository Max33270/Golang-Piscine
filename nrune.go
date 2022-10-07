package piscine

func NRune(s string, n int) rune {
	letter := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	} else {
		return letter[n-1]
	}
}
