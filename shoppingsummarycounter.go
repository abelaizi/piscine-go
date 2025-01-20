package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	sim := make(map[string]int)
	words := split(str, " ")
	for _, word := range words {
		sim[word]++
	}
	return sim
}

func split(s, sep string) []string {
	var result []string
	length := len(sep)
	current := 0
	for i := 0; i <= len(s)-length; i++ {
		if s[i:i+length] == sep {
			result = append(result, s[current:i])
			current = i + length
		}
	}
	result = append(result, s[current:])
	return result
}
