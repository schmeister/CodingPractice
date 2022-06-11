package linkedListCycle

import (
	"fmt"
	"testing"
)

func TestLinkedListCycle(t *testing.T) {

	head := ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
					},
				},
			},
		},
	}
	fmt.Println(head.LinkedListCycle())
	head.Next.Next.Next.Next = head.Next.Next
	fmt.Println(head.LinkedListCycle())

}
