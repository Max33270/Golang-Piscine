package piscine

func BasicAtoi2(s string) int {
	sChangeable := []byte(s)
	result := 0
	grandeur := 1
	for i := len(sChangeable) - 1; i >= 0; i-- {
		if int(sChangeable[i]) >= 48 && int(sChangeable[i]) < 58 {
			result += (int(sChangeable[i]) - 48) * grandeur
			grandeur *= 10
		} else {
			return 0
		}
	}
	return int(result)
}
