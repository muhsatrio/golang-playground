package array

func ReverseArray(a []int32) []int32 {
	// Write your code here
	results := make([]int32, len(a))
	idx := 0

	for i := len(a) - 1; i >= 0; i-- {
		results[idx] = a[i]
		idx++
	}

	return results
}
