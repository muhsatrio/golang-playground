// problem link: https://leetcode.com/problems/merge-two-sorted-lists/

package linkedlist

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	for list1 != nil || list2 != nil {
		if list1 == nil {
			result = Insert(result, list2.Val)
			list2 = list2.Next
		} else if list2 == nil {
			result = Insert(result, list1.Val)
			list2 = list1.Next
		} else {
			if list1.Val < list2.Val {
				result = Insert(result, list1.Val)
				list2 = list1.Next
			} else {
				result = Insert(result, list2.Val)
				list2 = list2.Next
			}
		}
	}
	return result
}
