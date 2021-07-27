// still not AC
package array

import (
	"sort"
)

func ArrayManipulation(n int32, queries [][]int32) int64 {

	arr := make([][]int32, 0)

	for _, eachQuery := range queries {
		arr = append(arr, []int32{eachQuery[0], eachQuery[2]})
		arr = append(arr, []int32{eachQuery[1] + 1, -1 * eachQuery[2]})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	var sum int32 = 0
	var max int32 = 0

	for _, eachArr := range arr {
		sum += eachArr[1]
		if sum > max {
			max = sum
		}
	}

	return int64(max)
}
