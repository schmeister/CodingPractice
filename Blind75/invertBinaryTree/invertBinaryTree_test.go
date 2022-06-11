package invertBinaryTree

import (
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	root := Node{
		Val: 4,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val:   1,
			},
			Right: &Node{
				Val:   3,
			},
		},
		Right: &Node{
			Val: 7,
			Left: &Node{
				Val:   6,
			},
			Right: &Node{
				Val:   9,
			},
		},
	}
	print(&root, 0, 0, 0)
	invertTree(&root)
	print(&root, 0, 0, 0)
}
