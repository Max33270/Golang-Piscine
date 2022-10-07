package piscine

func AppendRange(min int, max int) []int {
	var emptyArray []int
	if min >= max {
		return emptyArray
	} else if min < max {
		for i := min; i < max; i++ {
			emptyArray = append(emptyArray, i)
		}
	}
	return emptyArray
}
