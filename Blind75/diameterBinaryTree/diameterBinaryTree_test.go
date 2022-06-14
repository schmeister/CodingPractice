package diameterBinaryTree

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	root := Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
		},
	}
	answer := 0
	fmt.Println(root.DFS(&answer))
}
