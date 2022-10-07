package piscine

func ConcatParams(args []string) string {
	concatenation := ""
	for i := 0; i < len(args); i++ {
		if i == len(args)-1 {
			concatenation = concatenation + args[i]
		} else if i < len(args)-1 {
			concatenation = concatenation + args[i] + "\n"
		}
	}
	return concatenation
}
