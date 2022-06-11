package reverseLinkedList

import (
	"testing"
)

func TestReverseList(t *testing.T) {

	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
					},
				},
			},
		},
	}
	head.Traverse()
	reversed := ReverseList(&head)
	reversed.Traverse()
}
