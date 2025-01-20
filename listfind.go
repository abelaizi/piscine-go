package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	ras := l.Head

	for ras != nil {
		if comp(ras.Data, ref) {
			return &ras.Data
		}
		ras = ras.Next
	}

	return nil
}
