package maxProfit

import (
	"fmt"
	"testing"
)

// Peak Valley Approach
func TestMaxProfit(t *testing.T) {
	arr := []int{7, 1, 5, 3, 6, 4}
	arr2 := []int{7, 6, 4, 3, 1}

	fmt.Println(MaxProfit(arr))
	fmt.Println(MaxProfit(arr2))
}
