package lowestCommonAncestor

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func LeastCommonAncestorInBST(root *TreeNode, n1 int, n2 int) int {
	if root == nil {
		return 10000000
	}

	if n1 > root.Val && n2 > root.Val {
		return LeastCommonAncestorInBST(root.Right, n1, n2)
	}

	if n1 < root.Val && n2 < root.Val {
		return LeastCommonAncestorInBST(root.Left, n1, n2)
	}

	return root.Val
}

func Print(node *TreeNode, ns int, ch int, total int) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v", ch, node.Val)
	fmt.Printf("\n")

	Print(node.Left, ns+2, 'L', total+node.Val)
	Print(node.Right, ns+2, 'R', total+node.Val)
}