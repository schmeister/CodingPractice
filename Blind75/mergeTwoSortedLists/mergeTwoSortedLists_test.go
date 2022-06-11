package mergeTwoSortedLists

import "testing"


func TestMergeTwoSortedLists(t *testing.T) {

	head1 := &Node{
		val:  0,
		next: &Node{
			val:  2,
			next: &Node{
				val:  4,
				next: nil,
			},
		},
	}
	head2 := &Node{
		val:  1,
		next: &Node{
			val:  3,
			next: &Node{
				val:  5,
				next: nil,
			},
		},
	}
	head := MergeTwoSortedLists(head1, head2)
	head.Print()
}
