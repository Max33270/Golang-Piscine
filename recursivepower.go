package piscine

func RecursivePower(nb int, power int) int {
	nb1 := nb
	if nb > 1000 {
		return 0
	} else if power == 0 || power == 0 && nb == 0 {
		return 1
	} else {
		if power < 0 || nb == 0 {
			return 0
		} else {
			nb1 *= RecursivePower(nb, power-1)
			return nb1
		}
	}
}
