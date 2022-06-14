package kClosestPoints

import (
	"fmt"
	"testing"
)

func TestKClosest(t *testing.T) {
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1

	fmt.Println(KClosest(points, k))
	fmt.Println(KClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
}
