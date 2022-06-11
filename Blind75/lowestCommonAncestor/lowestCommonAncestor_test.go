package lowestCommonAncestor

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {

	root := TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}

	Print(&root, 0, 0, 0)
	fmt.Println(LeastCommonAncestorInBST(&root, 5, 9))
	fmt.Println(LeastCommonAncestorInBST(&root, 2, 4))
}
