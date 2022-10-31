package list

// type ListInterfase interface {
// 	Head
// }

// Returns index of first element with given value
func Find(l *LinkedList, value interface{}) (int, *node) {
	n := l.Head
	idx := 0

	for n != nil {
		if n.Value == value {
			return idx, n
		} else {
			n = n.Next
			idx++
		}
	}

	return -1, nil
}
