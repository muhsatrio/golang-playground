// problem link: https://leetcode.com/problems/add-two-numbers/

package linkedlist

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var value, savedValue int
	var result *ListNode
	for l1 != nil || l2 != nil {
		upValue := 0
		downValue := 0
		if l1 != nil {
			upValue = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			downValue = l2.Val
			l2 = l2.Next
		}
		totalValue := upValue + downValue + savedValue
		if totalValue < 10 {
			value = totalValue
			savedValue = 0
		} else {
			value = totalValue % 10
			savedValue = totalValue / 10
		}

		// insert node
		result = Insert(result, value)
	}
	if savedValue > 0 {
		result = Insert(result, savedValue)
	}
	return result
}
