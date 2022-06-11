package mergeTwoSortedLists

import "fmt"

type Node struct {
	val  int
	next *Node
}

func MergeTwoSortedLists(head1 *Node, head2 *Node) *Node {

	cur := &Node{
		val:  -1,
		next: nil,
	}
	var head = cur

	for {
		if head1 != nil && head2 != nil {
			if cur == nil && head1.val <= head2.val {
				cur = head1
				head = head1
				head1 = head1.next
			} else if cur == nil && head1.val > head2.val {
				cur = head
				head = head2
				head2 = head2.next
			} else if head1.val <= head2.val {
				cur.next = head1
				head1 = head1.next
			} else {
				cur.next = head2
				head2 = head2.next
			}
			cur = cur.next
		} else {
			break
		}
	}
	if head1 != nil {
		cur.next = head1
	}
	if head2 != nil {
		cur.next = head2
	}

	return head.next
}

func (head *Node) Print() {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}
