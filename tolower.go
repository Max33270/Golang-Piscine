package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 65 && runes[i] <= 90 {
			runes[i] = (rune(runes[i]) + 32)
		}
	}
	return string(runes)
}
