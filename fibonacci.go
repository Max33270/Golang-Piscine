package piscine

func Fibonacci(index int) int {
	result := 0
	if index < 0 {
		return -1
	}
	if index == 0 {
		result = 0
		return 0
	}
	if index == 1 || index == 2 {
		result = 1
		return 1
	}
	if index == 0 {
		result = 0
	}
	if index >= 0 {
		result = Fibonacci(index-1) + Fibonacci(index-2)
		return result
	}
	return 0
}
