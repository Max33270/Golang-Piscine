package piscine

func AlphaCount(s string) int {
	counter := 0
	letter := []rune(s)
	for i := 0; i < len(s); i++ {
		if 65 <= (rune(letter[i])) && (rune(letter[i])) <= 90 || 97 <= (rune(letter[i])) && (rune(letter[i])) <= 122 {
			counter += 1
		}
	}
	return counter
}
