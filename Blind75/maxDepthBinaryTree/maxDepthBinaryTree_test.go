package maxDepthBinaryTree

import (
	"fmt"
	"testing"
)

func TestMaxDepthBinaryTree(t *testing.T) {

	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  &TreeNode{
				Val:   0,
				Left:  &TreeNode{
					Val:   0,
				},
			},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 17,
			},
		},
	}

	fmt.Println(root.MaxDepthBinaryTree())
}
