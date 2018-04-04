package main

type node struct {
	data int
	next *node
}

func newNode(d int) *node {
	return &node{
		data: d,
	}
}

// LinkedList ...
type LinkedList struct {
	root *node
}

// AddItem ...
func (l *LinkedList) AddItem(d int) {
	if l.root == nil {
		l.root = newNode(d)
		return
	}

	cur := l.root
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode(d)
}

// RemoveItem ...
func (l *LinkedList) RemoveItem(d int) {
	if l.root == nil {
		return
	}

	var prev *node
	cur := l.root
	for cur != nil {
		if d == cur.data {
			// cur is first item so update root
			if prev == nil {
				l.root = cur.next
			} else {
				prev.next = cur.next
			}
			cur = nil
			break
		}
		prev = cur
		cur = cur.next
	}
}
