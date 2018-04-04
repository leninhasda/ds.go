package main

import "github.com/davecgh/go-spew/spew"

func main() {
	list := &LinkedList{}
	list.AddItem(1)
	list.AddItem(2)
	list.AddItem(3)
	list.AddItem(4)
	list.AddItem(5)

	list.RemoveItem(3)
	list.AddItem(3)
	spew.Dump(list)
}
