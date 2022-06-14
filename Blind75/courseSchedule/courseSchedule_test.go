package courseSchedule

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	fmt.Println(CanFinish(2, [][]int{{1, 0}}))
//	fmt.Println(CanFinish(2, [][]int{{3, 2}}))
	fmt.Println(CanFinish(4, [][]int{{0, 1},{1, 2},{1, 3},{2, 3},{3, 0}}))
//	fmt.Println(CanFinish(4, [][]int{{0, 1},{1, 0},{2, 3}}))
}
