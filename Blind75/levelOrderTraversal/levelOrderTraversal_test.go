package levelOrderTraversal

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tree := TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
				Val:   15,
			},
			Right: &TreeNode{
				Val:   7,
			},
		},
	}
	fmt.Println(LevelOrder(&tree))
}
