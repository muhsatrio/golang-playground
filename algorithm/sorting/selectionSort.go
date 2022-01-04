package sorting

func SelectionSort(lists []int) []int {
	n := len(lists)
	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if lists[j] < lists[minIdx] {
				minIdx = j
			}
		}

		// swap

		temp := lists[i]
		lists[i] = lists[minIdx]
		lists[minIdx] = temp
	}

	return lists
}
