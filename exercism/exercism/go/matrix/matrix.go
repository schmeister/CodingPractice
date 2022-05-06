package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	lenR := len(rows)
	matrix := make(Matrix, lenR)

	for idxR, row := range rows {
		row = strings.TrimSpace(row)
		cols := strings.Split(row, " ")
		lenC := len(cols)
		matrix[idxR] = make([]int, lenC)

		for idxC, cell := range cols {

			iVal, err := strconv.Atoi(cell)
			if err != nil {
				return nil, err
			}
			matrix[idxR][idxC] = iVal
		}
		if idxR > 0 && len(matrix[idxR]) != len(matrix[idxR-1]) {
			return nil, errors.New("Bad size")
		}
	}
	return &matrix, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	if len(*m) == 0 {
		return [][]int{}
	}
	arr := make([][]int, 0)
	for j := 0; j < len((*m)[0]); j++ {
		r := make([]int, 0)
		for i := 0; i < len(*m); i++ {
			r = append(r, (*m)[i][j])
		}
		arr = append(arr, r)
	}
	return arr
}

func (m *Matrix) Rows() [][]int {
	arr := make([][]int, len(*m))
	for i := 0; i < len(*m); i++ {
		r := make([]int, len((*m)[0]))
		copy(r, (*m)[i])
		arr[i] = r
	}
	return arr
}

func (m *Matrix) Set(row, col, val int) bool {
	if len(*m) == 0 || row >= len(*m) || col >= len((*m)[0]) || row < 0 || col < 0 {
		return false
	}
	(*m)[row][col] = val
	return true
}
