package linkedlist

import (
	"fmt"
	"log"
)

type Node struct {
	next  *Node
	value interface{}
}

type SingleLinkedList struct {
	head *Node
	last *Node
}

func (L *SingleLinkedList) Insert(value interface{}) {
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

func (L *SingleLinkedList) Remove() {
	if L.head == L.last {
		L.head = nil
		L.last = nil
	} else {
		list := L.head
		for list.next != L.last {
			list = list.next
		}
		L.last = list
		L.last.next = nil
	}
}

func (L *SingleLinkedList) RemoveAfter(key interface{}) {
	if L.head == L.last {
		L.head = nil
		L.last = nil
	} else {
		list := L.head
		for list != L.last && list.value != key {
			list = list.next
		}
		if list.next == nil {
			log.Fatalln("node is empty")
		}
		P := list.next
		list.next = P.next
		P.next = nil
	}
}

func (L *SingleLinkedList) RemoveHead() {
	if L.head == L.last {
		L.head = nil
		L.last = nil
	} else {
		P := L.head
		L.head = P.next
		P.next = nil
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

func (L *SingleLinkedList) InsertAfter(key interface{}, value interface{}) {
	list := L.head
	for list.value != key && list.next != nil {
		list = list.next
	}
	if list.value != key {
		log.Fatalln("node target is not found")
	}
	newNode := &Node{
		value: value,
	}
	if list == L.last {
		newNode.next = list.next
		L.last = newNode
	} else {
		newNode.next = list.next
	}
	list.next = newNode
}

func (L *SingleLinkedList) PrintAll() {
	if L.head == nil {
		log.Fatalln("node is empty")
	}
	list := L.head
	for list.next != nil {
		fmt.Println(list.value)
		list = list.next
	}
	fmt.Println(list.value)
}
