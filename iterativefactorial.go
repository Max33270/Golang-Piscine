package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb < 0 || nb >= 1000 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb < 1000 {
		for i := 0; i < nb+1; i++ {
			if i == 1 {
				result = 1
			} else {
				result *= i
			}
		}
	}
	return result
}
