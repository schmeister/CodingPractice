package searchRotatedSortedArray

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {

	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(Search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(Search([]int{1}, 0))
}
