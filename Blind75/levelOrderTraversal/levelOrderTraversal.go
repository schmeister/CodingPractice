package levelOrderTraversal

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}

	// Key: Level, Value: TreeNode.Val
	m := map[int][]int{}

	root.DFS(m, 0)
	fmt.Println(m)
	return ret
}

func (root *TreeNode) DFS(m map[int][]int, level int) {
	if root == nil {
		return
	}

	nodes := m[level]
	nodes = append(nodes, root.Val)
	m[level] = nodes

	root.Left.DFS(m, level+1)
	root.Right.DFS(m, level+1)
}
