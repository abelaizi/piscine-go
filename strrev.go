package piscine

func StrRev(s string) string {
	var o string
	for i := len(s) - 1; i >= 0; i-- {
		o = o + string(s[i])
	}
	return o
}
