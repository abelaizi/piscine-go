package piscine

func AlphaCount(s string) int {
	counter := 0
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			counter += 1
		}
	}
	return counter
}
