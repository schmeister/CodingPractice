package tree

import "fmt"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	dict := make(map[int]*Node)

	// Iterate over incoming records
	fmt.Printf("Incoming\n")
	for _, r := range records {
		// See if dict contains this record
		if _, ok := dict[r.ID]; ok {
			// It does, do nothing
		} else {
			// it does not, create it, and add to dict
			dict[r.ID] = &Node{
				ID:       r.ID,
				Children: []*Node{},
			}
		}
		fmt.Printf("ID:%d Parent:%d\n", r.ID, r.Parent)
	}
	fmt.Printf("On to Map\n")

	for _, val := range dict{
		fmt.Printf("ID:%d Parent:%d\n", val.ID, len(val.Children))
	}
	return &Node{}, nil
}
