package array

func MatchingStrings(strings []string, queries []string) []int32 {
	// Write your code here
	results := make([]int32, 0)
	for _, eachQuery := range queries {
		var count int32 = 0
		for _, eachString := range strings {
			if eachQuery == eachString {
				count++
			}
		}
		results = append(results, count)
	}
	return results
}
