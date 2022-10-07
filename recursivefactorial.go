package piscine

func RecursiveFactorial(nb int) int {
	result := nb
	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else if nb >= 500 {
		return 0
	} else if nb < 500 {
		if nb == 1 {
			return 1
		} else if nb == 2 {
			return 2
		} else {
			result *= RecursiveFactorial(nb - 1)
		}
	}
	return result
}
