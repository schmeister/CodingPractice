package middleLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) FindMiddle() *ListNode {
	if node == nil {
		return nil
	}

	curr := node
	fast := node.Next

	for fast != nil {
		curr = curr.Next
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
	}
	return curr
}
