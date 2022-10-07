package piscine

func BasicAtoi(s string) int {
	sChangeable := []byte(s)
	result := 0
	grandeur := 1
	for i := len(sChangeable) - 1; i >= 0; i-- {
		result += (int(sChangeable[i]) - 48) * grandeur
		grandeur *= 10
	}
	return int(result)
}
