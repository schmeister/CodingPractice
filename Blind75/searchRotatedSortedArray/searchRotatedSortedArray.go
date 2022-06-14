package searchRotatedSortedArray

func Search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	if start == end {
		return -1
	}

	for {
		if start > end {
			break
		}

		// Locate mid-point - is it a match?
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}

		// Now search either the (remaining) upper or (remaining) lower portion 
		// Start is lower than mid
		if nums[start] <= nums[mid] {
			// Target is in this section. Set end to Mid -1 and re-search 
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				// Target in the upper section
				start = mid + 1
			}
		} else {
			// Rotation is in the upper section.
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
