package binarySearch

import (
	"fmt"
	"testing"
)


func TestDFS(t *testing.T) {

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
	fmt.Println(root.DFS(6))
	fmt.Println(root.DFS(65))
	fmt.Println(root.DFS(5))
	fmt.Println(root.DFS(4))
}
