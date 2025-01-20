package piscine

func Index(s string, toFind string) int {
	for j := 0; j <= len(s)-len(toFind); j++ {
		if s[j:j+len(toFind)] == toFind {
			return j
		}
	}
	return -1
}
