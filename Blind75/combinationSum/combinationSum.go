package combinationSum

import (
	"fmt"
	"sort"
)

func CombinationSumUp(candidates []int, target int) {
	size := len(candidates) - 1
	for i := size; i >= 0; i-- {
		k := i
		t := target
		group := make([]int, 0)
		pos := make([]int, 0)
		sum := 0

		for sum < t && k >= 0 {
			curr := candidates[k]
			sum += curr

			if sum <= target {
				group = append(group, candidates[k])
				pos = append(pos, k)
			} else {
				k--
				sum -= curr
			}
			if sum == target {
				fmt.Println("--->", group, sum)
				rem := group[len(group)-1]
				sum -= rem

				group = group[:len(group)-1]
				kk := pos[len(pos)-1]
				pos = pos[:len(pos)-1]
				k = kk - 1
			}
		}
	}
}

func sumArray(arr []int) int {
	sum := 0
	for _,val := range arr {
		sum += val
	}
	return sum
}

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ret := make([][]int, 0)
	ans := make([]int, 0)
	dfs(candidates, target, 0, &ans, &ret)
	return ret
}

func dfs(nums []int, target, i int, ans *[]int, ret *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ret = append(*ret, append(make([]int, 0), (*ans)...))
		return
	}
	for j := i; j < len(nums) && target >= nums[j]; j++ {
		*ans = append(*ans, nums[j])
		dfs(nums, target-nums[j], j, ans, ret)
		*ans = (*ans)[:len(*ans)-1]
	}
}