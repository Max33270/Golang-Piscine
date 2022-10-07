package piscine

func IterativePower(nb int, power int) int {
	nb1 := 1
	if nb > 1000 {
		return 0
	} else if power == 0 || power == 0 && nb == 0 {
		return 1
	} else {
		if power < 0 || nb == 0 {
			return 0
		} else {
			for i := 0; i < power; i++ {
				nb1 = nb1 * nb
			}
			return nb1
		}
	}
}
