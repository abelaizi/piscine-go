package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Tail = n
		l.Head = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
}
