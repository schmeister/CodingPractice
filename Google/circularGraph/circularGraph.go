package circularGraph

import "fmt"

type Node struct {
	key   int
	child map[int]*Node
}

var NodeMap = make(map[int]*Node)

func newNode(key int) *Node {
	var me *Node = &Node{}
	me.key = key
	me.child = make(map[int]*Node, 0)
	return me
}

func addEdge(x, y int) {
	// Find or create Node
	from, ok := NodeMap[x]
	if !ok {
		from = newNode(x)
		NodeMap[x] = from
	}

	to, ok := NodeMap[y]
	if !ok {
		to = newNode(y)
		NodeMap[y] = to
	}

	from.child[y] = to
}

func findCycle(node *Node, visited map[int]bool) bool {
	fmt.Printf("%d ", node.key)
	if ok, _ := visited[node.key]; ok {
		fmt.Print("cycle")
		return true
	}
	visited[node.key] = true
	if node == nil {
		return false
	}

 	for i, temp := range node.child {
		cycle := findCycle(temp, visited)
		if cycle {
			return true
		}
		i++
	}
	// Recurse return - not found
	visited[node.key] = false
	return false
}

func CircularGraph() {
	addEdge(1, 2)
	addEdge(1, 3)
	addEdge(1, 4)
	addEdge(1, 4)
	addEdge(2, 3)
	addEdge(3, 4)
	addEdge(4, 1)
	
	for _, y := range NodeMap {
		m := make(map[int]bool)
		findCycle(y, m)
		fmt.Println()
	}
}
