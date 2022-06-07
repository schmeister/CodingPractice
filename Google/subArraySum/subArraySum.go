package subArraySum

import "fmt"

func subArraySum(nums []int, k int) (int, int) {

	low := 0
	high := 0
	sum := nums[low]

	for low < (len(nums)) {
		if sum < k {
			high++
			if high >= len(nums) {
				return -1, -1
			}
			sum += nums[high]
		} else if sum > k {
			sum -= nums[low]
			low++
		} else if sum == k {
			return low, high
		}
	}

	return -1, -1
}

func SubArraySum() {

	arr := []int{3, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(subArraySum(arr, 18))

}
