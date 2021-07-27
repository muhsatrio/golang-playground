package main

import (
	"fmt"
	"golang-playground/hackerrank/datastructure/array"
)

func main() {
	var n, m int
	var a, b, k int32

	queries := make([][]int32, 0)

	fmt.Scan(&n)
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		fmt.Scan(&a)
		fmt.Scan(&b)
		fmt.Scan(&k)
		queries = append(queries, []int32{a, b, k})
	}

	fmt.Println(array.ArrayManipulation(int32(n), queries))

}
