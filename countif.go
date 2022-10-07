package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for i := range tab {
		if f(tab[i]) == true {
			counter++
		}
	}
	return counter
}
