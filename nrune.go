package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}
	runes := []rune(s)
	return runes[n-1]
}
