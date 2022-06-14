package matrix01

import (
	"fmt"
	"testing"
)

func TestMatrix01(t *testing.T) {

	fmt.Println(UpdateMatrix([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 0}}))
	fmt.Println(UpdateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(UpdateMatrix([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))
	fmt.Println(UpdateMatrix([][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}
