package piscine

func DescendAppendRange(max, min int) []int {
	sliceRange := []int{}
	if min < max {
		for i := max; i > min; i-- {
			sliceRange = append(sliceRange, i)
		}
		return sliceRange
	}
	return []int{}
}
