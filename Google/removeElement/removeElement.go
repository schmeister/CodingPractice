package removeElement

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (list *List) add(val int) {
	if list.head == nil {
		list.head = &Node{}
		list.head.data = val
	} else {
		list.head.add(val)
	}
}

func (node *Node) add(val int) {
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{
		data: val,
	}
}

func (list *List) print() {
	if list.head != nil {
		node := list.head
		for node != nil {
			fmt.Println(node.data)
			node = node.next
		}
	}
}

func (list *List) remove(val int) {
	if list.head != nil {
		curr := list.head
		if curr.data == val {
			list.head = curr.next
		} else {
			for curr.next != nil {
				next := curr.next
				if next.data == val{
					curr.next = next.next
				}
				curr = curr.next
			}
		}

	}
}

func RemoveElement() {
	list := &List{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)
	list.add(5)
	list.add(6)
	list.add(7)
	list.add(8)
	list.add(9)

	list.print()
	list.remove(2)
	list.remove(6)
	list.remove(2)
	list.print()
}
