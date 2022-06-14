package numIslands

import "fmt"

func NumIslands(grid [][]int) int {
	var count int
	for x := range grid {
		for y := range grid[x] {
			// Check every cell if part of an island
			if grid[x][y] == 1 {
				fmt.Println(grid)
				search(grid, x, y)
				count++
			}
		}
	}
	return count
}

// Search finds all positions to current island and sets them to 0
func search(grid [][]int, x, y int) {

	// Out of bounds.
	xOut := x < 0 || x >= len(grid)
	yOut := y < 0 || y >= len(grid[0])

	// Return if not part of an island.
	if xOut || yOut || grid[x][y] == 0 {
		return
	}

	// Current location is part of island, now remove it from contention for future checks.
	grid[x][y] = 0

	// Check all adjacent cells
	search(grid, x-1, y)
	search(grid, x+1, y)
	search(grid, x, y-1)
	search(grid, x, y+1)
}
