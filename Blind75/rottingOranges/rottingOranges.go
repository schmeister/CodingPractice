package rottingOranges

import "fmt"

type Point struct {
	i, j int
}

func RottingOranges(grid [][]int) int {

	// Find coordinates of initial rotten oranges
	var coords []Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				coords = append(coords, Point{i, j})
			}
		}
	}
	fmt.Println(coords)

	// Each change/iteration from good to rotten is 1 minute.
	minutes := 0
	for len(coords) > 0 {
		s := len(coords)
		rotten := false
		// Change all adj. oranges to rotter for ever orange already rotten.
		for i := 0; i < s; i++ {
			// Pull first rotten orange off slice and turn adjacent all bad.
			p := coords[0]
			coords = coords[1:]

			// If adjacent cells have oranges, turn them to rotten.
			// Add rotten oranges to rotten coordinates.
			if p.i+1 < len(grid) && grid[p.i+1][p.j] == 1 {
				grid[p.i+1][p.j] = 2
				coords = append(coords, Point{p.i + 1, p.j})
				rotten = true
			}
			if p.i-1 >= 0 && grid[p.i-1][p.j] == 1 {
				grid[p.i-1][p.j] = 2
				coords = append(coords, Point{p.i - 1, p.j})
				rotten = true
			}
			if p.j+1 < len(grid[0]) && grid[p.i][p.j+1] == 1 {
				grid[p.i][p.j+1] = 2
				coords = append(coords, Point{p.i, p.j + 1})
				rotten = true
			}
			if p.j-1 >= 0 && grid[p.i][p.j-1] == 1 {
				grid[p.i][p.j-1] = 2
				coords = append(coords, Point{p.i, p.j - 1})
				rotten = true
			}
		}
		if rotten {
			minutes++
		}
	}

	// Check to see if there are any good oranges left.
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return minutes
}
