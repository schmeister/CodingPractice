package trees

import (
	"fmt"
	"strings"
)

var sb strings.Builder

type Node struct {
	data     int
	children []*Node
	visited  bool
}

func show(data []*Node) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("["))
	for idx, n := range data {
		sb.WriteString(fmt.Sprintf("%d", n.data))
		if idx < (len(data) - 1) {
			sb.WriteString(fmt.Sprintf(" "))
		}
	}
	sb.WriteString(fmt.Sprintf("]"))
	return sb.String()
}

func (root *Node) addChildByValue(i int) Node {
	n := newNode(i)
	root.addChildByNode(&n)
	return n
}

func (root *Node) addChildByNode(n *Node) {
	root.children = append(root.children, n)
}

func newNode(val int) Node {
	n := Node{
		data:     val,
		children: make([]*Node, 0),
		visited:  false,
	}
	return n
}

func (root *Node) preOrder() {
	if root == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%d ", root.data))

	for _, node := range root.children {
		node.preOrder()
	}
}

func Nary() {
	root := &Node{
		data:     10,
		children: []*Node{},
	}
	root.addChildByValue(8)
	root.addChildByValue(5)

	root.children[0].addChildByValue(2)
	root.children[0].addChildByValue(1)
	root.children[0].addChildByValue(6)

	// Add child node [9,11] in node (1)
	root.children[0].children[1].addChildByValue(9)
	root.children[0].children[1].addChildByValue(11)
	// Add child node [7  8 3  4] in node (5)
	root.children[1].addChildByValue(7)
	root.children[1].addChildByValue(8)
	root.children[1].addChildByValue(3)
	root.children[1].addChildByValue(4)
	// Add child node [-7] in node (4)
	root.children[1].children[0].addChildByValue(-1)
	// Add child node [2,1,3] in node (7)
	root.children[1].children[3].addChildByValue(2)
	root.children[1].children[3].addChildByValue(1)
	root.children[1].children[3].addChildByValue(3)
	root.preOrder()
	fmt.Printf("%s -> %5s\n", "N-ary Tree - PreOrder", sb.String())
}

func BFS() {
	zero := newNode(0)
	one := newNode(1)
	two := newNode(2)
	three := newNode(3)

	zero.addChildByNode(&one)
	zero.addChildByNode(&two)
	one.addChildByNode(&two)
	two.addChildByNode(&zero)
	two.addChildByNode(&three)
	three.addChildByNode(&three)

	fmt.Printf("%s\n", two.bfs())
}

func (root *Node) bfs() string {
	nodes := make([]*Node, 0)
	nodes = append(nodes, root)

	i := 0
	l := len(nodes)
	for i < l {
		node := nodes[i]
		node.visited = true

		for _, n := range node.children {
			if !n.visited && !contains(nodes, n) {
				nodes = append(nodes, n)
			}
		}

		l = len(nodes)
		i++
	}
	return fmt.Sprintf("BFS: %s", show(nodes))
}

func contains(nodes []*Node, item *Node) bool {
	found := false
	for _, n := range nodes {
		if n == item {
			return true
		}
	}
	return found
}
