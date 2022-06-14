package validateBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsValidBST(root *TreeNode) bool {
	MaxUint := ^uint(0)
	MaxInt := int(MaxUint >> 1)
	MinInt := -MaxInt - 1
	return root.isValidbst(MinInt, MaxInt)
}

func (root *TreeNode) isValidbst(min, max int) bool {
	if root == nil {
		return true
	}

	left := root.Left.isValidbst(min, root.Val)
	right := root.Right.isValidbst(root.Val, max)

	return root.Val < max && root.Val > min && left && right
}
