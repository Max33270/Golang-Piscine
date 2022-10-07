package piscine

func StrRev(s string) string {
	rstring := ""
	for i := len(s) - 1; i >= 0; i-- {
		rstring += string(s[i])
	}
	return rstring
}
