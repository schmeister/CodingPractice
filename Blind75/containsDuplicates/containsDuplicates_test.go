package containsDuplicates

import (
	"fmt"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {

	fmt.Println(ContainsDuplicates([]int{1, 2, 3, 1}))
	fmt.Println(ContainsDuplicates([]int{1, 2, 3, 4}))
	fmt.Println(ContainsDuplicates([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
