package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] < 65 || runes[i] > 90 {
			return false
		}
	}
	return true
}
