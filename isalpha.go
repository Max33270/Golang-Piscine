package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 48 || runes[i] > 57 && runes[i] < 65 || runes[i] > 90 && runes[i] < 97 || runes[i] > 122 {
			return false
		}
	}
	return true
}
