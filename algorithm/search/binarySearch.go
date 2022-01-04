package search

func BinarySearch(list []int, item int) bool {
	left := 0
	right := len(list) - 1
	for left < right {
		mid := (left + right) / 2

		if item < list[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if list[left] == item {
		return true
	} else {
		return false
	}
}
