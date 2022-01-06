package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func Insert(node *ListNode, value int) *ListNode {
	temp := &ListNode{
		Val: value,
	}
	if node == nil {
		node = temp
	} else {
		nextNode := node
		for nextNode.Next != nil {
			nextNode = nextNode.Next
		}
		nextNode.Next = temp
	}
	return node
}
