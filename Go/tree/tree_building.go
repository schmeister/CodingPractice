package tree

import (
	"errors"
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	Parent   int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	dict := make(map[int]*Node)

	// Iterate over incoming records and add to map
	fmt.Printf("\n")
	for _, r := range records {
		fmt.Printf("Received: (Parent)%d --> (ID)%d\n", r.Parent, r.ID)
		node := Node{
			ID:       r.ID,
			Children: make([]*Node, 0),
			Parent:   r.Parent}

		_, ok := dict[node.ID]
		if ok {
			fmt.Printf("DUPLICATE %d\n", node.ID)
			return dict[0], errors.New("Duplicate")
		}
		dict[node.ID] = &node
	}
	fmt.Printf("dict size: %d\n", len(dict))

	visited := make(map[int]bool)
	for _, node := range dict {

		if node.Parent == node.ID || visited[node.ID] {
			continue
		}
		visited[node.ID] = true

		myParent := node.Parent
		pNode := dict[myParent]

		if pNode != nil && node.Parent <= node.ID {
			dict[myParent].Children = append(dict[myParent].Children, node)

			sort.SliceStable(dict[myParent].Children, func(i, j int) bool {
				return dict[myParent].Children[i].ID < dict[myParent].Children[j].ID
			})
		} else if node.Parent > node.ID {
			return nil, errors.New("Root node has parent")
		}
	}

	fmt.Printf("DONE\n")
	vm = make(map[int]bool)

	value, ok := dict[0]
	if !ok {
		fmt.Printf("root node: \n")
		return value, errors.New("No root node")
	} else {
		PrintMap(dict)
		Preorder(dict[0], "")
		return value, nil
	}
}

func PrintMap(dict map[int]*Node) {
	for _, n := range dict {
		children := n.Children
		fmt.Printf("%d --> (", n.ID)
		for idx, child := range children {
			fmt.Printf("%d", child.ID)
			if idx < len(children)-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf(")\n")
	}
}

var vm map[int]bool

func Preorder(root *Node, space string) {
	_, ok := vm[root.ID]
	if ok {
		return
	} else {
		//	vm[root.ID] = true
	}

	fmt.Printf("%s%d\n", space, root.ID)
	for _, c := range root.Children {
		Preorder(c, space+"  ")
	}
	//	panic("")
}
