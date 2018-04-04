package main

import "fmt"

type dnode struct {
	data int
	next *dnode
	prev *dnode
}

func newDnode(d int) *dnode {
	return &dnode{
		data: d,
	}
}

// DLinkedList ...
type DLinkedList struct {
	root *dnode
}

// AddItem ...
func (l *DLinkedList) AddItem(d int) {
	if l.root == nil {
		l.root = newDnode(d)
		return
	}

	cur := l.root
	for cur.next != nil {
		cur = cur.next
	}
	dn := newDnode(d)
	dn.prev = cur
	cur.next = dn
}

// RemoveItem ...
func (l *DLinkedList) RemoveItem(d int) {
	if l.root == nil {
		return
	}

	cur := l.root
	for cur != nil {
		if d == cur.data {
			if cur.prev != nil {
				cur.prev.next = cur.next
			} else {
				// cur is first item so update root
				l.root = cur.next
			}
			if cur.next != nil {
				cur.next.prev = cur.prev
			}
			cur = nil
			break
		}
		cur = cur.next
	}
}

// Print ...
func (l *DLinkedList) Print() {
	if l.root == nil {
		return
	}
	cur := l.root
	l.print(cur)
	for cur.next != nil {
		cur = cur.next
		l.print(cur)
	}
}

func (l *DLinkedList) print(n *dnode) {
	if n.prev != nil {
		fmt.Print("prev: ", n.prev.data)
	} else {
		fmt.Print("prev: nil")
	}
	fmt.Print(" data: ", n.data)
	if n.next != nil {
		fmt.Print(" next: ", n.next.data)
	} else {
		fmt.Print(" next nil")
	}
	fmt.Println("\n--------------------------")
}
