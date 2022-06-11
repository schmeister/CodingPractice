package twoSum

func TwoSum(arr []int, goal int) (int, int) {

	left := 0
	right := len(arr) - 1

	for left < right {
		l := arr[left]
		r := arr[right]

		sum := l + r

		if sum < goal {
			left++
		} else if sum > goal {
			right--
		} else if sum == goal {
			return l, r
		}
	}
	return -1, -1
}
