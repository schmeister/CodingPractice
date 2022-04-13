package trees

import "fmt"

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

func (root *Node) preOrder() {
	if root == nil {
		return
	}
	fmt.Println(root.data)

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

	root.preOrder()
}