package array

func DynamicArray(n int32, queries [][]int32) []int32 {
	// Write your code here
	result := make([]int32, 0)
	arr := make([][]int32, n)
	var lastAnswer int32 = 0
	for _, eachElement := range queries {
		if eachElement[0] == 1 {
			idx := ((eachElement[1] ^ lastAnswer) % n)
			arr[idx] = append(arr[idx], eachElement[2])
		} else {
			idx := ((eachElement[1] ^ lastAnswer) % n)
			lastAnswer = arr[idx][eachElement[2]%int32(len(arr[idx]))]
			result = append(result, lastAnswer)
		}
	}
	return result
}
