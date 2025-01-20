package piscine

func TrimAtoi(s string) int {
	sign := 1
	re := 0

	for i, str := range s {
		if i == 0 && (str == '-' || str == '+') {
			if str == '-' {
				sign = -1
			}
			continue
		} else if (i != 0) && re == 0 && str == '-' {
			sign = -1
		}

		if str < '0' || str > '9' {
			continue
		} else {
			z := int(str - '0')
			re = re*10 + z
		}
	}
	return re * sign
}
