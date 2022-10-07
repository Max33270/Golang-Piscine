package piscine

func MakeRange(min, max int) []int {
	counter := 0
	if max-min <= 0 {
		return nil
	} else if max-min >= 0 {
		answer := make([]int, max-min)
		for i := min - 1; i < max-1; i++ {
			answer[counter] = i + 1
			counter++
		}
		return answer
	}
	return nil
}
