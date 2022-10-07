package piscine

func Split(s, sep string) []string {
	var res []string
	var nextSplit int

	for len(s) > 0 {
		nextSplit = Index(s, sep)
		if nextSplit != -1 {
			res = append(res, s[0:nextSplit])
			s = s[nextSplit+len(sep):]
		} else {
			res = append(res, s)
			s = ""
		}
	}
	return res
}
