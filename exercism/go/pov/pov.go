package pov

type Tree struct {
	value    string
	parent   *Tree
	children []*Tree
}

func New(value string, children ...*Tree) *Tree {
	root := &Tree{
		value:    value,
		children: children,
		parent:   nil,
	}
	for _, child := range children {
		child.parent = root
	}

	return root
}

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) String(lvl int) string {
	if tr == nil {
		return "nil"
	}
	result := tr.Value()
	if len(tr.Children()) == 0 {
		return result
	}
	for _, ch := range tr.Children() {
		result += " (" + ch.String(lvl+1) + ")"
	}
	return "(" + result + ")"
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	//	Find the new root node
	stack := make([]string, 0)
	node, found := tr.DFS_with_trace(from, &stack)
	if !found {
		return nil
	}

	seen := make(map[string]bool)
	ntree := Rotate(node, seen)
	return ntree
}

// PathTo returns the shortest path between two nodes in the tree.
func (tr *Tree) PathTo(from, to string) []string {
	// Get the path trace from the tr node to the "to" and "from" nodes
	fromStack := make([]string, 0)
	toStack := make([]string, 0)
	_, fromFound := tr.DFS_with_trace(from, &fromStack)
	_, toFound := tr.DFS_with_trace(to, &toStack)
	if !fromFound || !toFound {
		return []string{}
	}

	// Find the shorter path, and locate where they diverge.
	minIDX := len(toStack)
	if len(fromStack) < len(toStack) {
		minIDX = len(fromStack)
	}
	i := 0
	for ; i < minIDX; i++ {
		if fromStack[i] != toStack[i] {
			break
		}
	}

	// Formulate the path trace and return
	path := make([]string, 0)
	for x := len(fromStack) - 1; x >= (i - 1); x-- {
		path = append(path, fromStack[x])
	}
	for x := i; x < len(toStack); x++ {
		path = append(path, toStack[x])
	}
	return path
}

/////////////////////

func Rotate(tr *Tree, seen map[string]bool) *Tree {

	if tr == nil || seen[tr.value] {
		return nil
	}
	seen[tr.value] = true

	children := make([]*Tree, 0, len(tr.children)+1)
	// Rotate Children
	for _, child := range tr.children {
		cTree := Rotate(child, seen)
		if cTree != nil {
			children = append(children, cTree)
		}
	}

	// Rotate parent down
	if tr.parent != nil {
		pTree := Rotate(tr.parent, seen)
		if pTree != nil {
			children = append(children, pTree)
		}
	}
	// Return tr as root
	return New(tr.value, children...)
}

// DFS - Depth-first traversal to designated node.
func (tr *Tree) DFS_with_trace(value string, stack *[]string) (*Tree, bool) {
	*stack = append(*stack, tr.value)

	// Node found, return node with trace.
	if tr.value == value {
		return tr, true
	}

	// Use recusion to search children until found
	found := false
	for _, child := range tr.children {
		tr, found = child.DFS_with_trace(value, stack)
		if found {
			break
		}
	}

	// Will need to remove last element
	l := len(*stack)
	if l > 0 && !found {
		s := *stack
		s = s[:l-1]
		*stack = s
	}
	return tr, found
}
