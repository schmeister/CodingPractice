package numIslands

import (
	"fmt"
	"testing"
)



func TestNumIslands(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	grid2 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}

	fmt.Println(NumIslands(grid))
	fmt.Println(NumIslands(grid2))
}
