package piscine

func ReverseMenuIndex(menu []string) []string {
	arr := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		arr[i] = menu[len(menu)-1-i]
	}
	return arr
}
