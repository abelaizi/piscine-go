package piscine

func Unmatch(a []int) int {
	var s int
	for _, e := range a {
		s = 0
		for _, v := range a {
			if v == e {
				s++
			}
		}
		if s%2 != 0 {
			return e
		}
	}
	return -1
}
