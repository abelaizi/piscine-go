package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	rsl := strs[0]
	for i := 1; i < len(strs); i++ {
		rsl += sep + strs[i]
	}
	return rsl
}
