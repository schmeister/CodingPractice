package diameterBinaryTree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (root *Node) DFS(depth *int) int {
	if root == nil {
		return 0
	}

	// Get the maximum value for each path
	ld := root.Left.DFS(depth)
	rd := root.Right.DFS(depth)

	// Return the maximum values
	if *depth > (ld + rd) {
		return *depth
	}
	if ld > rd {
		return ld + 1
	}
	return rd + 1
}
