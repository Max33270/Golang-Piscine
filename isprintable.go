package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if runes[i] >= 32 && runes[i] <= 127 {
		} else {
			return false
		}
	}
	return true
}
