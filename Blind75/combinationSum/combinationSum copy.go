package combinationSum

import (
	"fmt"
)

func CombinationSumUp_org(candidates []int, target int) {
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
