package piscine

func JumpOver(str string) string {
	s := ""
	if len(str) >= 3 {
		for i := 2; i < len(str); i += 3 {
			s += string(str[i])
		}
	}
	s += "\n"
	return s
}
