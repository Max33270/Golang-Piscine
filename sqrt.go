package piscine

func Sqrt(nb int) int {
	var result int
	for i := 0; i < nb+1; i++ {
		result = nb * nb
		if result == nb {
			return nb
		}
	}
	return 0
}
