package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 97 && runes[i] <= 122 {
			runes[i] = (rune(runes[i]) - 32)
		}
	}
	return string(runes)
}
