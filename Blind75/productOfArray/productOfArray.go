package productOfArray

func ProductOfArray(nums []int) []int {
	length := len(nums)
	result := make([]int, length)

	left := make([]int,length)
	right := make([]int, length)
	left[0], right[length-1] = 1, 1

	// Work from the outside in (but not self)
	// Multiply adjacent values
	for i, j := 1, length-2; i < length && j >= 0; i, j = i+1, j-1 {
		left[i] = left[i-1] * nums[i-1]
		right[j] = right[j+1] * nums[j+1]
	}

	// Multiply L * R
	for i := range result {
		result[i] = left[i] * right[i]
	}
	return result
}
