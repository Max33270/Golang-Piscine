package piscine

func TrimAtoi(s string) int {
	var res string
	for _, item := range s {
		if len(res) == 0 && item == 45 {
			res += string('-')
		} else if IsNumeric(string(item)) {
			res += string(item)
		}
	}
	return Atoi(res)
}
