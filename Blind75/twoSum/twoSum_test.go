package twoSum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	vals := []int{2, 4, 6, 8, 10}
	fmt.Println(vals)
	fmt.Println(TwoSum(vals, 6))
	fmt.Println(TwoSum(vals, 14))
}
