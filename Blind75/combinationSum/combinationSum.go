package combinationSum

// CombinationSum2 - Bad performance O(N^2)
func CombinationSum2(candidates []int, target int) [][]int {

	size := len(candidates) - 1
	result := make([][]int, 0)
	for i := size; i >= 0; i-- {
		k := i
		t := target
		group := make([]int, 0)
		pos := make([]int, 0)
		sum := 0

		for sum < t && k >= 0 {
			curr := candidates[k]
			group = append(group, curr)
			pos = append(pos, k)
			sum = sumArray(group)

			if sum == target {
				result = append(result, append(make([]int, 0), (group)...))
				group = group[:len(group)-1]
				k = pos[len(pos)-1] - 1
				pos = pos[:len(pos)-1]
			} else if sum > target {
				group = group[:len(group)-1]
				k = pos[len(pos)-1] - 1
				pos = pos[:len(pos)-1]
			}

			for k < 0 && len(pos) > 1 {
				group = group[:len(group)-1]
				k = pos[len(pos)-1] - 1
				pos = pos[:len(pos)-1]
				sum = sumArray(group)
			}
			sum = sumArray(group)
		}
	}
	return result
}

func sumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func CombinationSum(candidates []int, target int) [][]int {
	ret := make([][]int, 0)
	ans := make([]int, 0)
	dfs(candidates, target, 0, &ans, &ret)
	return ret
}

func dfs(nums []int, remainder, idx int, ans *[]int, ret *[][]int) {
	if remainder < 0 {
		return
	}
	if remainder == 0 {
		*ret = append(*ret, append(make([]int, 0), (*ans)...))
		return
	}
	for j := idx; j < len(nums) && remainder >= nums[j]; j++ {
		*ans = append(*ans, nums[j])
		dfs(nums, remainder-nums[j], j, ans, ret)
		*ans = (*ans)[:len(*ans)-1]
	}
}
