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

	cur := l.root
	pre := l.root
	for cur.next != nil {
		if d == cur.data {
			pre.next = cur.next
			break
		}
		pre = cur
		cur = cur.next
	}
}
