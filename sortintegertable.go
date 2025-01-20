package piscine

func SortIntegerTable(table []int) {
	var temp int
	for i := 0; i <= len(table)-1; i++ {
		for m := i + 1; m <= len(table)-1; m++ {
			if table[i] > table[m] {
				temp = table[i]
				table[i] = table[m]
				table[m] = temp
			}
		}
	}
}
