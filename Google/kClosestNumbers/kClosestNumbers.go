package kClosestNumbers

import "fmt"

// https://shareablecode.com/snippets/golang-solution-for-leetcode-problem-find-k-closest-elements-R4TQ-fkWN

func findClosestNumbers(arr []int, count int, target int) []int {
	min := 0
	max := len(arr) - count
	for min < max {
		mid := min + (max-min)>>1
		if target-arr[mid] > arr[mid+count]-target {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return arr[min : min+count]
}

func FindClosestElement() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	target := 7
	count := 5

	fmt.Printf("%v\n", findClosestNumbers(arr, count, target))
}
