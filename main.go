package main

import (
	"fmt"
	"golang-playground/algorithm/search"
	"golang-playground/algorithm/sorting"
	"golang-playground/leetcode/array"
	"golang-playground/leetcode/linkedlist"
)

func main() {
	fmt.Println(array.RemoveElement([]int{3, 2, 2, 3}, 3))

	// implementation binary search

	fmt.Println(search.BinarySearch([]int{1, 2, 3}, 5))

	// implementation selection sort

	fmt.Println(sorting.SelectionSort([]int{2, 3, 9, 7}))

	// implementation buuble sort

	fmt.Println(sorting.SelectionSort([]int{2, 10, 9, 7}))

	var angka1, angka2 *linkedlist.ListNode

	angka1 = linkedlist.Insert(angka1, 2)
	angka1 = linkedlist.Insert(angka1, 4)
	angka1 = linkedlist.Insert(angka1, 3)

	angka2 = linkedlist.Insert(angka2, 5)
	angka2 = linkedlist.Insert(angka2, 6)
	angka2 = linkedlist.Insert(angka2, 4)

	_ = linkedlist.AddTwoNumbers(angka1, angka2)

}
