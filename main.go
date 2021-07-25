package main

import (
	"fmt"
	"golang-playground/hackerrank/datastructure/array"
)

func main() {
	var n int
	var temp int32

	var a []int32

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		a = append(a, temp)
	}

	result := array.ReverseArray(a)

	for _, eachElement := range result {
		fmt.Println(eachElement)
	}
}
