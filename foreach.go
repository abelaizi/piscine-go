package piscine

func ForEach(f func(int), a []int) {
	for _, rest := range a {
		f(rest)
	}
}
