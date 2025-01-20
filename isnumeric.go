package piscine

func IsNumeric(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
