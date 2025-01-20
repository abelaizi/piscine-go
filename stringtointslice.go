package piscine

func StringToIntSlice(str string) []int {
	var mySlice []int

	for _, c := range str {
		mySlice = append(mySlice, int(c))
	}
	return mySlice
}
