package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 48 || runes[i] > 57 {
			return false
		}
	}
	return true
}
