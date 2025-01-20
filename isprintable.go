package piscine

func IsPrintable(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 32 && s[i] <= 127 {
			continue
		} else {
			return false
		}
	}
	return true
}
