package instertInterval

import (
	"fmt"
	"testing"
)

func TestInstertInterval(t *testing.T) {

	fmt.Println(InsertInterval([][]int{{1, 3}, {6, 9}}, []int{4, 5}))
	fmt.Println(InsertInterval([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	fmt.Println(InsertInterval([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	fmt.Println(InsertInterval([][]int{{1, 2}, {4, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{3, 11}))
}
