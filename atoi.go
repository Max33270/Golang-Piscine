package piscine

func Atoi(s string) int {
	var resultat int = 0
	var compte int = 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 43 && i == 0 {
			return resultat
		}
		if s[i] == 45 && i == 0 {
			resultat *= -1
			return resultat
		}
		if s[i] < 48 || s[i] > 57 {
			return 0
		}
		resultat += int(s[i]-48) * compte
		compte = 10 * compte
	}
	return resultat
}
