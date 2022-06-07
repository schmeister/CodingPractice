package pathsForASum

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (tree *Tree) Insert(val int) {
	if tree.Root == nil {
		tree.Root = &Node{
			Data:  val,
			Left:  nil,
			Right: nil,
		}
	} else {
		root := tree.Root
		root.Insert(val)
	}
}

func (node *Node) Insert(val int) {
	if node == nil {
		return
	} else if val <= node.Data {
		if node.Left == nil {
			node.Left = &Node{
				Data: val,
			}
		} else {
			node.Left.Insert(val)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{
				Data: val,
			}
		} else {
			node.Right.Insert(val)
		}
	}
}

func print(w io.Writer, node *Node, ns int, ch rune, goal int, total int) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v", ch, node.Data)
	if goal == (total + node.Data) {
		fmt.Fprintf(w, " - Success %d", goal)
	}
	fmt.Fprintf(w, "\n")

	print(w, node.Left, ns+2, 'L', goal, total+node.Data)
	print(w, node.Right, ns+2, 'R', goal, total+node.Data)
}

func PathsForASum() {
	tree := Tree{}
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(1)
	tree.Insert(1)
	tree.Insert(-1)
	tree.Insert(4)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(2)

	print(os.Stdout, tree.Root, 0, 'M', 5, 0)
}
