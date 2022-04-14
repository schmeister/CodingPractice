package arrays

import "fmt"

func subarraySum(slice []int, goal int) (int, int) {

	len := len(slice)

	for i := 0; i < len; i++ {
		sum := slice[i]
		for j := i + 1; j < len; j++ {
			sum += slice[j]
			if sum > goal {
				break
			}
			if sum == goal {
				return i, j
			}
		}
	}
	return 0, 0
}

func SubArraySum() {
	slice := []int{1, 4, 20, 3, 10, 5}
	goal := 33
	l, r := subarraySum(slice, goal)
	fmt.Printf("%-15s%5d%5d = %d\n", "SubArraySum",goal,l, r)
}
