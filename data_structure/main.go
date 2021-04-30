package main

import (
	"golang-playground/data_structure/linkedlist"
)

func main() {
	linkedList := &linkedlist.SingleLinkedList{}
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.PrintAll()
}
