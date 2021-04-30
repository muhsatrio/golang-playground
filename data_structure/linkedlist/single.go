package linkedlist

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type SingleLinkedList struct {
	head *Node
	last *Node
}

func (L *SingleLinkedList) Add(value interface{}) {
	newNode := &Node{
		value: value,
	}
	if L.head == nil && L.last == nil {
		L.head = newNode
		L.last = newNode
	} else {
		L.last.next = newNode
		L.last = newNode
	}
}

func (L *SingleLinkedList) InsertHead(value interface{}) {
	newNode := &Node{
		value: value,
	}
	if L.head == nil && L.last == nil {
		L.head = newNode
		L.last = newNode
	} else {
		newNode.next = L.head
		L.head = newNode
	}
}

func (L *SingleLinkedList) PrintAll() {
	list := L.head
	for list.next != nil {
		fmt.Println(list.value)
		list = list.next
	}
	fmt.Println(list.value)
}
