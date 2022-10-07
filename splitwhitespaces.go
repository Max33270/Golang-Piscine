package piscine

func SplitWhiteSpaces(s string) []string {
	counter := 0 // number of separators (+1) corresponds to number of words, separators are spaces, tabs and lines
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '	' || s[i] == '\n' { // if it's a separator
			if i+1 < len(s) && i != 0 { // i cannot be 1st of last char
				if s[i+1] != ' ' && s[i+1] != '	' && s[i+1] != '\n' { // the char following can be a separator
					counter++ // total number of separators
				}
			}
		}
	}
	emptyArray := make([]string, counter+1) // creates table of strings where strings depend on number spaces
	index := 0                              // to be in within the correct string
	for i := 0; i < len(s); i++ {           // same loop as above increases instead of separator count
		if s[i] == ' ' || s[i] == '	' || s[i] == '\n' {
			if i+1 < len(s) && i != 0 {
				if s[i+1] != ' ' && s[i+1] != '	' && s[i] != '\n' {
					index++
				}
			}
		} else {
			emptyArray[index] = emptyArray[index] + string(s[i])
		}
	}
	return emptyArray
}
