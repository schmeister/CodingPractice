package arbitraryPointer

import "fmt"

type Node struct {
	Arb  *Node
	Next *Node
	Key  int
}

// Append adds a node to the end of the list.
func (node *Node) Append(key int) {
	new := &Node{
		Next: nil,
		Arb:  nil,
		Key:  key,
	}

	n := node
	for n.Next != nil {
		n = n.Next
	}
	n.Next = new

	n = node
	for n.Arb != nil {
		n = n.Arb
	}
	n.Arb = new
}

// Duplicate inserts a duplicate element between original and Next
func (head *Node) Duplicate() *Node {
	curr := head

	// Duplicate and insert between nodes
	for curr != nil {
		new := Node{
			Key:  curr.Key + 10,
			Next: curr.Next,
		}
		curr.Next = &new
		curr = new.Next
	}

	// Assign Abitrary nodes - Current Node to Arbitrary.Next
	curr = head
	head2 := head.Next
	for curr != nil {
		if curr.Arb != nil && curr.Arb.Next != nil {
			curr.Next.Arb = curr.Arb.Next
		}
		// move to the next newly added node by
		// skipping an original node
		curr = curr.Next.Next
	}
	// Return the head of the duplicated list
	return head2
}

// Remove the duplicate inserted elements and create new list
func (head *Node) Separate() {
	// Break the LinkedList into two lists
	curr := head
	dup := head.Next
	for curr.Next != nil && dup.Next != nil {
		curr.Next = curr.Next.Next
		dup.Next = dup.Next.Next
		curr = curr.Next
		dup = dup.Next
	}
	curr.Next = nil
	dup.Next = nil
}

func (n *Node) Display() {
	fmt.Printf("(next)")
	for n != nil {
		fmt.Printf("%+v", n.Key)
		if n.Next != nil {
			fmt.Printf(" -> ")
		}
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	fmt.Println()
}

func (n *Node) DisplayArb() {
	fmt.Printf("(arb) ")
	for n != nil {
		fmt.Printf("%+v", n.Key)
		if n.Arb != nil {
			fmt.Printf(" -> ")
		}
		if n.Arb == nil {
			break
		}
		n = n.Arb
	}
	fmt.Println()
}

func ArbitraryPointer() {

	for x := 2; x < 5; x++ {
		head := &Node{
			Next: nil,
			Key:  0,
		}
		for y := 1; y <= x; y++ {
			head.Append(y)
		}

		fmt.Println("Initial==============================")
		head.Display()
		head.DisplayArb()
		/*
			fmt.Println("Duplicated==============================")
			head.Display()
			dup.Display()
			head.DisplayArb()
			dup.DisplayArb()
		*/
		dup := head.Duplicate()
		head.Separate()
		fmt.Println("Separate==============================")
		head.Display()
		dup.Display()
		head.DisplayArb()
		dup.DisplayArb()
	}
}
