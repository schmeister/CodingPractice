package binarySearch

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (root *Node) DFS(find int) bool {
	if root == nil {
		return false
	}
	if root.Val == find{
		return true
	}
	found := root.Left.DFS(find)
	if !found {
		found = root.Right.DFS(find)
	}
	return found
}
