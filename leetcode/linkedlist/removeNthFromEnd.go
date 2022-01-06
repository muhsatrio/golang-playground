package linkedlist

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	lengthNode := 1
	tempNode := head
	for tempNode.Next != nil {
		lengthNode++
		tempNode = tempNode.Next
	}

	if lengthNode == 1 && n == 1 {
		return nil
	}

	var previousNode *ListNode

	i := 0
	tempNode = head
	for i < lengthNode-n {
		previousNode = tempNode
		tempNode = tempNode.Next
		i++
	}

	if previousNode == nil {
		head = head.Next
	} else if tempNode.Next == nil {
		previousNode.Next = nil
	} else {
		previousNode.Next = tempNode.Next
	}
	return head
}
