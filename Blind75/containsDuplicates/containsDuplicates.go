package containsDuplicates

func ContainsDuplicates(arr []int) bool {

	seen := make(map[int]bool)
	for _, i := range arr {
		if seen[i] {
			return true
		}
		seen[i] = true
	}
	return false
}
