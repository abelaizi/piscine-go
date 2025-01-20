package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		return l.Tail.Data
	}
}
