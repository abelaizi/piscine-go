package piscine

func IsUpper(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
