package rottingOranges

import (
	"fmt"
	"testing"
)

func TestRottingOranges(t *testing.T) {

	grid1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
//	grid2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
//	grid3 := [][]int{{0, 2}}

	fmt.Println(RottingOranges(grid1))
//	fmt.Println(RottingOranges(grid2))
//	fmt.Println(RottingOranges(grid3))
}
