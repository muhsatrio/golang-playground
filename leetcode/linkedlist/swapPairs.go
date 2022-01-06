// problem link: https://leetcode.com/problems/swap-nodes-in-pairs/

package linkedlist

func SwapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var previousNode, tempNode, tempPreviousNode, tempNextNode, result *ListNode

	node := head

	swap := false

	for node != nil {
		if swap {
			tempNode = node
			tempPreviousNode = previousNode
			tempNextNode = node.Next
			tempNode.Next = nil
			tempPreviousNode.Next = nil
			result = Insert(result, tempNode.Val)
			result = Insert(result, tempPreviousNode.Val)
			tempNode.Next = tempPreviousNode
			tempPreviousNode.Next = tempNextNode
			node = tempNextNode
			previousNode = tempPreviousNode
			swap = false
		} else {
			if node.Next == nil {
				Insert(result, node.Val)
			}
			previousNode = node
			node = node.Next
			swap = true
		}
	}

	return result
}
