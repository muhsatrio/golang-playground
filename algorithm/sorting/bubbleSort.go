package sorting

func BubbleSort(lists []int) []int {
	n := len(lists)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if lists[j] > lists[j+1] {
				// swap

				temp := lists[j]
				lists[j] = lists[j+1]
				lists[j+1] = temp
			}
		}
	}

	return lists
}
