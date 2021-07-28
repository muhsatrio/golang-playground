package linkedlist

import "fmt"

func PrintLinkedList(head *SinglyLinkedListNode) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
