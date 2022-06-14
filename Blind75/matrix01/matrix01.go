package matrix01

import (
	"math"
)

///////////////
func UpdateMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	maxRow := len(matrix)
	maxCol := len(matrix[0])

	// Reset all non-zero cells to maximum
	numZeros := 0
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if matrix[r][c] != 0 && !AdjZero(matrix, r, c) {
				matrix[r][c] = math.MaxInt64
			} else if matrix[r][c] == 0 {
				numZeros++
			}
		}
	}
	// No Zero entries - no solution
	if numZeros == 0 {
		return matrix
	}

	// Treat each cell as the root of a tree
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if matrix[r][c] == 1 {
				DFS(matrix, r, c, -1)
			}
		}
	}
	return (matrix)
}

// AdjZero checks to see if there are adjacent cells that have a zero
func AdjZero(matrix [][]int, row, col int) bool {
	if row > 0 && matrix[row-1][col] == 0 {
		return true
	}
	if col > 0 && matrix[row][col-1] == 0 {
		return true
	}
	if row < len(matrix)-1 && matrix[row+1][col] == 0 {
		return true
	}
	if col < len(matrix[0])-1 && matrix[row][col+1] == 0 {
		return true
	}
	return false
}

// DFS for an n-ary "tree"
func DFS(matrix [][]int, row, col, val int) {
	// invalid cell
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) {
		return
	}

	// no need to update or already visited
	if (matrix[row][col] <= val) {
		return
	}
	
	// lower value found
	if val > 0 {
		matrix[row][col] = val
	}
	DFS(matrix, row-1, col, matrix[row][col]+1)
	DFS(matrix, row, col-1, matrix[row][col]+1)
	DFS(matrix, row+1, col, matrix[row][col]+1)
	DFS(matrix, row, col+1, matrix[row][col]+1)
}
