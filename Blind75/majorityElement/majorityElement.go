package majorityElement

import "fmt"

func MajorityElement(arr []int) int {
	m := make(map[int]int)
	elements := len(arr)

	for _, val := range arr {
		m[val]++
	}
	for key, val := range m {
		ratio := float64(val) / float64(elements)
		if ratio >= .5 {
			return key
		}
	}
	return 0
}
func MajorityElement2(nums []int) int {
	res := nums[0]
	count := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			res = nums[i]
		}
		if nums[i] == res {
			count++
		} else {
			count--
		}
	}
	ratio := float64(count) / float64(len(nums))
	fmt.Printf("%-5d%-5d%-5.2f\n",res, count, ratio)
	if ratio >= .5 {
		return res
	}

	return -1
}
