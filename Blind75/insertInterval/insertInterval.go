package instertInterval

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	start := newInterval[0]
	end := newInterval[1]
	result := make([][]int, 0)
	i := 0
	tempInterval := []int{start, end}

	// Add Intervals when start of new is after end of existing
	for ; i < len(intervals) && start > intervals[i][1]; i++ {
	}
	result = append(result, intervals[:i]...)

	// Collapse Intervals while end is after current start
	for ; i < len(intervals) && end >= intervals[i][0]; i++ {
		start = min(start, intervals[i][0])
		end = max(end, intervals[i][1])
		tempInterval = []int{start, end}
	}
	// Append Temp Interval
	result = append(result, tempInterval)
	// Append remaining Intervals
	result = append(result, intervals[i:]...)

	return result
}
func max(a, b int) int {
	if b > a {
		return b
	} else {
		return a
	}
}
func min(a, b int) int {
	if b < a {
		return b
	} else {
		return a
	}
}
