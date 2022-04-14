package trees

import (
	"fmt"
	"strings"
)

type Node struct {
	data     int
	children []*Node
}

func (root *Node) addChild(i int) {
	n := Node{
		data:     i,
		children: []*Node{},
	}
	root.children = append(root.children, &n)
}

var sb strings.Builder

func (root *Node) preOrder() {
	if root == nil {
		return
	}
	sb.WriteString(fmt.Sprintf("%d ",root.data))

	for _, node := range root.children {
		node.preOrder()
	}
}

func Nary() {
	root := &Node{
		data:     10,
		children: []*Node{},
	}
	root.addChild(8)
	root.addChild(5)

	root.children[0].addChild(2)
	root.children[0].addChild(1)
	root.children[0].addChild(6)

	// Add child node [9,11] in node (1)
    root.children[0].children[1].addChild(9)
    root.children[0].children[1].addChild(11)
    // Add child node [7  8 3  4] in node (5)
    root.children[1].addChild(7)
    root.children[1].addChild(8)
    root.children[1].addChild(3)
    root.children[1].addChild(4)
    // Add child node [-7] in node (4)
    root.children[1].children[0].addChild(-1)
    // Add child node [2,1,3] in node (7)
    root.children[1].children[3].addChild(2)
    root.children[1].children[3].addChild(1)
    root.children[1].children[3].addChild(3)
	root.preOrder()
	fmt.Printf("%s -> %5s\n", "N-ary Tree - PreOrder",sb.String())
}