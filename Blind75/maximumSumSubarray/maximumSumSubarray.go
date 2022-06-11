package maximumSumSubarray

func MaxSumSubarray(arr []int) int {
	UintSize := 32 << (^uint(0) >> 32 & 1)
	MaxInt := 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt := -MaxInt - 1

	currMax := 0
	max := MinInt

	for _, val := range arr {
		currMax += val
		if currMax > max {
			max = currMax
		}
		if currMax < 0 {
			currMax = 0
		}
	}
	return max
}
