package array

func RotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	results := make([]int32, 0)
	length := len(arr)
	loc := int(d) % length
	for i := loc; i < length; i++ {
		results = append(results, arr[i])
	}
	for i := 0; i < loc; i++ {
		results = append(results, arr[i])
	}
	return results
}
