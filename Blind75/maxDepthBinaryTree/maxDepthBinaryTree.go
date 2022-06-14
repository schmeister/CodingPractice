package maxDepthBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) MaxDepthBinaryTree() int {

	if node == nil {
		return 0
	}

	L := node.Left.MaxDepthBinaryTree()
	R := node.Right.MaxDepthBinaryTree()

	if L > R {
		return L + 1
	}
	return R + 1
}
