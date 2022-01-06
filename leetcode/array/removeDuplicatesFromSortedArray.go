// problem link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package array

func RemoveDuplicates(nums []int) int {
	i := 0
	if len(nums) == 0 {
		return 0
	}
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
