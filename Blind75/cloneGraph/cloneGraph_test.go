package cloneGraph

import (
	"fmt"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	AddEdge(1, 2)
	AddEdge(1, 3)
	AddEdge(1, 4)
	AddEdge(1, 4)
	AddEdge(2, 3)
	AddEdge(3, 4)
	AddEdge(4, 1)

	for _, y := range NodeMap {
		m := make(map[int]bool)
		FindCycle(y, m)
		fmt.Println()
	}

	clone := CloneGraph(NodeMap[1])
	FindCycle(clone, make(map[int]bool))

}
