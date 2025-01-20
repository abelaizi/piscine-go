package piscine

func AppendRange(min, max int) []int {
	var integer []int
	for i := min; i < max; i++ {
		integer = append(integer, i)
	}
	return integer
}
