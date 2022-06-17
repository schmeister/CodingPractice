package searchRotatedSortedArray

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(Search([]int{4, 5, 6, 7, 8, 9, 0, 1, 2}, 1))
	fmt.Println(Search([]int{8, 9, 0, 1, 2, 3, 4, 5, 6}, 9))
	fmt.Println(Search([]int{1}, 0))
}
