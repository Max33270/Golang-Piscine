package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	counter := 1
	var result string
	for i := 0; i < len(s); i++ {
		if (runes[i] < 97 && runes[i] > 90) || (runes[i] < 48 && runes[i] > 0) || (runes[i] < 65 && runes[i] > 57) || runes[i] > 122 {
			counter += 1
			result += string(runes[i])
		} else {
			if counter > 0 {
				if runes[i] >= 97 && runes[i] <= 122 {
					runes[i] -= 32
				}
				counter = 0
				result += string(runes[i])
			} else {
				if runes[i] > 64 && runes[i] < 91 {
					runes[i] += 32
				}
				result += string(runes[i])
			}
		}
	}
	return result
}
