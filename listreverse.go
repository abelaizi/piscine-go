package piscine

func ListReverse(l *List) {
	var rev *NodeL
	for l.Head != nil {
		next := l.Head.Next
		l.Head.Next = rev
		rev = l.Head
		l.Head = next
	}
	l.Head = rev
}
