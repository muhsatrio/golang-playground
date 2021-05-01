package main

import (
	"golang-playground/data_structure/linkedlist"
)

func main() {
	linkedList := &linkedlist.SingleLinkedList{}
	linkedList.Insert(2)
	linkedList.Insert(3)
	linkedList.InsertAfter(3, 10)
	linkedList.Insert(4)
	linkedList.InsertHead(1)
	linkedList.Remove()
	linkedList.RemoveAfter(10)
	linkedList.PrintAll()
}
