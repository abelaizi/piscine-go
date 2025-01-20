package piscine

func SplitWhiteSpaces(s string) []string {
	var r []string
	wrd := ""

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if wrd != "" {
				r = append(r, wrd)
				wrd = ""
			}
		} else {
			wrd += string(char)
		}
	}

	if wrd != "" {
		r = append(r, wrd)
	}

	return r
}
