package piscine

func Abort(a, b, c, d, e int) int {
	r := []int{a, b, c, d, e}
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return r[2]
}
