package linkedLists

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func (parent *Node) addNode(d int) {
	child := Node{
		data: d,
		next: nil,
	}

	for parent.next != nil {
		parent = parent.next
	}
	parent.next = &child
}

func (parent *Node) walkList() string {
	var sb strings.Builder
	for parent != nil {
		sb.WriteString(fmt.Sprintf("%3d", parent.data))
		parent = parent.next
	}
	return sb.String()
}

func reverseKGroup(head *Node, k int) *Node {
	var tail *Node
	cur := head

	tempHead := &Node{0, head}
	previous := tempHead

	for cur != nil {
		tail = cur
		linkedListIndex := 0

		for cur != nil && linkedListIndex < k {
			cur = cur.next
			linkedListIndex++
		}

		if linkedListIndex < k {
			previous.next = tail
		} else {
			previous.next = reverse(tail, k)
			previous = tail
		}
	}
	return tempHead.next
}

//Create a function to easily reverse the linkedList
func reverse(head *Node, k int) *Node {

	var nextNode, previous *Node

	cur := head

	for cur != nil && k > 0 {
		nextNode = cur.next
		cur.next = previous
		previous = cur
		cur = nextNode
		k--
	}
	head = previous
	return head

}

// https://golangbyexample.com/reverse-nodes-k-group-linked-list-golang/
// https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/1087104/golang-reverse-node-in-k-group-short-and-easy-solution
func ReverseInGroups() {
	head := Node{
		data: 0,
		next: nil,
	}
	for i := 1; i < 10; i++ {
		head.addNode(i)
	}

	fmt.Println("ReverseInGroups - pre:  ", head.walkList())
	n := reverseKGroup(&head, 3)
	fmt.Println("ReverseInGroups - post: ", n.walkList())
}
