package firstBadVersion

func FirstBadVersion(numVersions int, firstBadVersion int) int {

	VAL = firstBadVersion

	max := numVersions
	min := 1

	for min < max {
		mid := (min + max) / 2
		if isBadVersion(mid) {
			max = mid-1
		} else {
			min = mid + 1
		}
	}
	return min
}

func FirstBadVersion2(n int, firstBadVersion int) int {
	VAL = firstBadVersion

	lo, hi := 1, n
	for lo < hi {
	  mid := (lo + hi) / 2
	  if isBadVersion(mid) {
		hi = mid
	  } else {
		lo = mid + 1
	  }
	}
	return lo
  }
var VAL int

func isBadVersion(val int) bool {
	return val <= VAL
}
