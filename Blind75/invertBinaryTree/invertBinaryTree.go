package invertBinaryTree

import "fmt"


type Node struct{
	Val int
	Left *Node
	Right *Node
}

func invertTree(root *Node) *Node {
    if root != nil {
        root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    }
    return root
}

func print(node *Node, ns int, ch int, total int) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v", ch, node.Val)

	fmt.Printf("\n")

	print(node.Left, ns+2, 'L', total+node.Val)
	print(node.Right, ns+2, 'R', total+node.Val)
}