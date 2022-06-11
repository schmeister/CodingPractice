package balancedTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)

	diff := abs(left - right)
	return diff <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
