package main

import "github.com/davecgh/go-spew/spew"

func linkedList() {
	list := &LinkedList{}
	list.AddItem(1)
	list.AddItem(2)
	list.AddItem(3)
	list.AddItem(4)
	list.AddItem(5)

	list.RemoveItem(3)
	list.RemoveItem(1)
	list.RemoveItem(5)
	list.AddItem(3)
	spew.Dump(list)
}

func dlinkedList() {
	list := &DLinkedList{}
	list.AddItem(1)
	list.AddItem(2)
	list.AddItem(3)
	list.AddItem(4)
	list.AddItem(5)
	list.RemoveItem(1)
	list.RemoveItem(2)
	list.RemoveItem(5)
	list.AddItem(6)

	list.Print()
}

func main() {
	dlinkedList()
}
