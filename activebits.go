package piscine

func ActiveBits(n int) int {
	count := 0
	for n > 0 {
		count += n % 2
		n = n / 2
	}
	return count
}
