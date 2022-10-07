package piscine

func LastRune(s string) rune {
	letter := []rune(s)
	return letter[len(s)-1]
}
