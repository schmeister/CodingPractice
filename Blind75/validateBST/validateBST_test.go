package validateBST

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	root := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(IsValidBST(&root))
	root2 := TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   0,
			},
			Right: &TreeNode{},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(IsValidBST(&root2))
}
