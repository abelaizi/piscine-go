package piscine

func ShoppingListSort(slice []string) []string {
	for f := 0; f < len(slice)-1; f++ {
		for g := f + 1; g < len(slice); g++ {
			if len(slice[f]) > len(slice[g]) {
				slice[f], slice[g] = slice[g], slice[f]
			}
		}
	}
	return slice
}
