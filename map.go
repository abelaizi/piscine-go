package piscine

func Map(f func(int) bool, a []int) []bool {
	i := []bool{}
	for _, s := range a {
		i = append(i, f(s))
	}
	return i
}
