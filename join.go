package piscine

func Join(strs []string, sep string) string {
	output := ""
	for i, value := range strs {
		output += value
		if len(strs)-1 > i {
			output += sep
		}
	}
	return output
}
