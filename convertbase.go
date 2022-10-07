package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	a := AtoiBase(nbr, baseFrom)
	b := PrintNbrBase(a, baseTo)
	return b
}

func AtoiBase(s string, base string) int {
	nbr := 0
	a := []rune(s)
	b := []rune(base)
	var index []int
	if !IsValidBase(base) {
		return 0
	}
	if IsValidBase(base) {
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(base); j++ {
				if a[i] == b[j] {
					index = append(index, j)
				}
			}
		}
		power := len(s) - 1
		for i := 0; i < len(s); i++ {
			nbr += index[i] * IterativePower(len(base), power)
			power--
		}
	}
	return nbr
}

func IsValidBase(base string) bool {
	runeBase := []rune(base)
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(runeBase); i++ {
		for j := 0; j < len(runeBase); j++ {
			if i != j {
				if runeBase[i] == runeBase[j] {
					return false
				}
			}
		}
	}
	for k := 0; k < len(runeBase); k++ {
		if runeBase[k] == 45 || runeBase[k] == 43 {
			return false
		}
	}
	return true
}

func IterativePower2(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else {
		for i := 1; i < power+1; i++ {
			result = result * nb
		}
		return result
	}
}

func PrintNbrBase(nbr int, base string) string {
	var str []rune
	if !IsValidBase(base) {
		return ""
	}
	result := ""
	if IsValidBase(base) {
		var number []int
		if nbr < 0 {
			result = "-"
			nbr = -nbr
		}
		rest := 0
		for i := 0; nbr != 0; i++ {
			rest = nbr % len(base)
			nbr = nbr / len(base)
			number = append(number, rest)

		}
		for i := len(number) - 1; i >= 0; i-- {
			runeBase := []rune(base)
			if number[i] < 0 {
				number[i] = -number[i]
			}
			index := number[i]
			str = append(str, runeBase[index])
		}
	}
	for i := 0; i < len(str); i++ {
		result += string(str[i])
	}
	return result
}
