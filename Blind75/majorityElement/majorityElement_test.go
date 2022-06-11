package majorityElement

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {

	arr := []int{2, 2, 1, 1, 1, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	//	fmt.Println(MajorityElement(arr))
	fmt.Println(MajorityElement2(arr))
	arr2 := []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 5, 5, 5, 5}
	fmt.Println(MajorityElement2(arr2))
}
