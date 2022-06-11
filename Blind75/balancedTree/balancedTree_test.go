package balancedTree

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
				Val:   9,
				// Right: &TreeNode{
				// 	Val:   0,
				// 	Right: &TreeNode{
				// 		Val:   0,
				// 		Right: &TreeNode{
				// 			Val:   0,
				// 		},
				// 	},
				// },
			},
		},
	}

	Print(&root, 0, 0, 0)
	fmt.Println(IsBalanced(&root))
}

func Print(node *TreeNode, ns int, ch int, total int) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v", ch, node.Val)
	fmt.Printf("\n")

	Print(node.Left, ns+2, 'L', total+node.Val)
	Print(node.Right, ns+2, 'R', total+node.Val)
}
