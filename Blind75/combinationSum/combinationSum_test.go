package combinationSum

import (
	"fmt"
	"testing"
)

func TestCombinationSumUp(t *testing.T) {
	//	size := len(candidates) - 1
	CombinationSumUp([]int{2, 3, 4, 5}, 8)
	fmt.Println(CombinationSum([]int{2, 3, 4, 5}, 8))
}
