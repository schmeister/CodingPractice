package cloneGraph

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

var NodeMap = make(map[int]*Node)

func NewNode(val int) *Node {
	var me *Node = &Node{}
	me.Val = val
	me.Neighbors = make([]*Node, 0)
	return me
}

func AddEdge(x, y int) {
	// Find or create Node
	from, ok := NodeMap[x]
	if !ok {
		from = NewNode(x)
		NodeMap[x] = from
	}

	to, ok := NodeMap[y]
	if !ok {
		to = NewNode(y)
		NodeMap[y] = to
	}

	from.Neighbors = append(from.Neighbors, to)
}

func CloneGraph(node *Node) *Node {
	visited := make(map[int]*Node)
	dup := node.DFS2(visited)
	return dup
}

func (node *Node) DFS2(clones map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	clone := *node
	clone.Neighbors = make([]*Node, len(node.Neighbors))
	clones[clone.Val] = &clone
	for i, n := range node.Neighbors {
		if clones[n.Val] != nil {
			clone.Neighbors[i] = clones[n.Val]
		} else {
			clone.Neighbors[i] = n.DFS2(clones)
		}
	}
	return &clone
}

func FindCycle(node *Node, visited map[int]bool) bool {
	fmt.Printf("%d ", node.Val)
	if ok, _ := visited[node.Val]; ok {
		fmt.Print("cycle")
		return true
	}
	visited[node.Val] = true
	if node == nil {
		return false
	}

 	for i, temp := range node.Neighbors {
		cycle := FindCycle(temp, visited)
		if cycle {
			return true
		}
		i++
	}
	// Recurse return - not found
	visited[node.Val] = false
	return false
}