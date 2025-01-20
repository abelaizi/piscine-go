package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	cou := 0
	co := 0
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			cou++
		} else if f(a[i], a[i+1]) > 0 {
			co++
		} else {
			return true
		}
	}
	if co == len(a)-1 {
		return true
	} else if cou == len(a)-1 {
		return true
	}
	return false
}
