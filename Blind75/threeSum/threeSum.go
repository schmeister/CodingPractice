package threeSum

import (
	"sort"
)

func ThreeSum(arr []int) [][]int {
	res := [][]int{}
	sort.Ints(arr)
	count := len(arr) - 1

	// Loop until 2 down from max (for j and k)
	i := 0
	for i < count-1 {
		j := i + 1
		k := count

		sum := arr[i] + arr[j] + arr[k]

		// Increase middle element, but less than kth element
		for sum < 0 && j < (k-1) {
			j++
			sum = arr[i] + arr[j] + arr[k]
		}
		// Decrease top element, but greater than jth element
		for sum > 0 && k > (j+1) {
			k--
			sum = arr[i] + arr[j] + arr[k]
		}

		// Match, store this tuple
		if sum == 0 {
			res = append(res, []int{arr[i], arr[j], arr[k]})
		}
		i++
	}
	return res
}
