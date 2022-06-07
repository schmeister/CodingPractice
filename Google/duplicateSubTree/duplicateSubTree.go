package duplicateSubTree

import "fmt"

type Node struct {
	data  rune
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (tree *Tree) Insert(val rune) {
	if tree.root == nil {
		tree.root = &Node{
			data:  val,
			left:  nil,
			right: nil,
		}
	} else {
		tree.root.Insert(val)
	}
}

func (node *Node) Insert(val rune) {
	if node == nil {
		return
	} else if val == node.data {
		tn := &Node{
			data:  val,
			left:  nil,
			right: nil,
		}
		if node.left == nil {
			node.left = tn
		} else if node.right == nil {
			node.right = tn
		} else {
			node.left.Insert(val)
		}
	} else if val < node.data {
		if node.left == nil {
			node.left = &Node{
				data:  val,
				left:  nil,
				right: nil,
			}
		} else {
			node.left.Insert(val)
		}
	} else {
		if node.right == nil {
			node.right = &Node{
				data:  val,
				left:  nil,
				right: nil,
			}
		} else {
			node.right.Insert(val)
		}
	}
}

func print(node *Node, ns int, ch rune, total rune) {

	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v", ch, node.data)

	fmt.Printf("\n")

	print(node.left, ns+2, 'L', total+node.data)
	print(node.right, ns+2, 'R', total+node.data)
}

func postorder(p *Node, indent int) {
	if p != nil {
		if p.left != nil {
			postorder(p.left, indent+4)
		}
		if p.right != nil {
			postorder(p.right, indent+4)
		}
		if indent > 0 {
			fmt.Printf("%*s", indent, " ")
		}
		fmt.Printf("%d\n", p.data)
	}
}
func postorder2(p *Node, indent int) {
	if p != nil {
		if p.right != nil {
			postorder2(p.right, indent+2)
		}
		if indent > 0 {
			fmt.Printf("%*s", indent, "")
		}
		if p.right != nil {
			fmt.Printf(" /\n%*s", indent, "")
		}
		fmt.Printf("%c\n", p.data)
		if p.left != nil {
			fmt.Printf("%*s \\\n", indent, "")
			postorder2(p.left, indent+2)
		}
	}
}
func DuplicateSubTree() {
	tree := Tree{}
	tree.Insert('M')
	tree.Insert('M')
	tree.Insert('M')
	tree.Insert('G')
	tree.Insert('H')
	tree.Insert('C')
	tree.Insert('H')
	tree.Insert('G')
	tree.Insert('H')
	tree.Insert('C')
	tree.Insert('H')

	print(tree.root, 0, 'M', 0)
	//	postorder(tree.root, 0)
	postorder2(tree.root, 0)
}
