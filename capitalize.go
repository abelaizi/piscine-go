package piscine

func Capitalize(s string) string {
	st := ""
	upper := true
	for _, i := range s {
		if (i >= '0' && i <= '9') || (i >= 'A' && i <= 'Z') || (i >= 'a' && i <= 'z') {
			if upper {
				if i >= 'a' && i <= 'z' {
					st += string(i - 32)
				} else {
					st += string(i)
				}
			} else {
				if i >= 'A' && i <= 'Z' {
					st += string(i + 32)
				} else {
					st += string(i)
				}
			}
			upper = false
		} else {
			st += string(i)
			upper = true
		}
	}
	return st
}
