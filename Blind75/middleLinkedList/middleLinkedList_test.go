package middleLinkedList

import (
	"fmt"
	"testing"
)

func TestFindMiddle(t *testing.T) {
	head := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{
							Val:  6,
						},
					},
				},
			},
		},
	}
	fmt.Println(head.FindMiddle())
}
