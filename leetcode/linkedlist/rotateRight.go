package linkedlist

func RotateRight(head *ListNode, k int) *ListNode {
	var result, backNode *ListNode

	if head == nil {
		return head
	}

	lengthNode := 0
	tempNode := head
	for tempNode != nil {
		lengthNode++
		tempNode = tempNode.Next
	}
	if k >= lengthNode {
		k = k % lengthNode
	}
	if k > 0 {
		i := 0
		tempNode = head
		for i < (lengthNode - k) {
			tempNode = tempNode.Next
			i++
		}
		backNode = tempNode
		for tempNode != nil {
			result = Insert(result, tempNode.Val)
			tempNode = tempNode.Next
		}
		tempNode = head
		for tempNode != backNode {
			result = Insert(result, tempNode.Val)
			tempNode = tempNode.Next
		}
	} else {
		return head
	}
	return result
}
