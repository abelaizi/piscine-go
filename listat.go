package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos && l != nil; i++ {
		l = l.Next
	}
	return l
}
