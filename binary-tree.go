package main

// Adder is a interface wrapping the tree add nodes method
type Adder interface {
	Add(node NodeComponent) error
}

// NodeValueGetter is a interface wrapping get Value methods.
type NodeValueGetter interface {
	GetValue() int64
}

// NodeLeftGetter is a interface wrapping the get left leaf.
type NodeLeftGetter interface {
	GetLeft() NodeComponent
}

// NodeLeftSetter is a interface wrapping the set left leaf.
type NodeLeftSetter interface {
	SetLeft(node NodeComponent)
}

// NodeRightGetter is a interface wrapping the get right leaf.
type NodeRightGetter interface {
	GetRight() NodeComponent
}

// NodeRightSetter is a interface wrapping the set right leaf.
type NodeRightSetter interface {
	SetRight(node NodeComponent)
}

// NodeGetter is a interface wrapping node get methods.
type NodeGetter interface {
	NodeValueGetter
	NodeLeftGetter
	NodeRightGetter
}

// NodeSetter is a interface wrapping node set methods.
type NodeSetter interface {
	NodeLeftSetter
	NodeRightSetter
}

// NodeComponent is a interface wrapping methods used for a node.
type NodeComponent interface {
	NodeGetter
	NodeSetter
}

// NewNode creates a new empty node.
func NewNode(value int64) *Node {
	return &Node{Value: value}
}

// Node represents a leaf node on the binary tree.
type Node struct {
	Value       int64
	left, right NodeComponent
}

// GetValue returns node value.
func (n Node) GetValue() int64 {
	return n.Value
}

// GetLeft returns the left leaf.
func (n Node) GetLeft() NodeComponent {
	return n.left
}

// SetLeft set the left leaf.
func (n *Node) SetLeft(node NodeComponent) {
	n.left = node
}

// GetRight returns the right leaf.
func (n Node) GetRight() NodeComponent {
	return n.right
}

// SetRight set the right leaf.
func (n *Node) SetRight(node NodeComponent) {
	n.right = node
}

// bfwalk does a breadth first walk of the node tree populating the
// seen slice with nodes found in order right to left leaf nodes. `max`
// is how many levels deep to go. `level` is the current level and should
// be 0 to start with.
func bfwalk(node NodeComponent, max, level int, seen *[]NodeComponent) {
	if level < max {
		level++

		if node.GetRight() != nil {
			*seen = append(*seen, node.GetRight())
		}

		if node.GetLeft() != nil {
			*seen = append(*seen, node.GetLeft())
		}

		if node.GetRight() != nil {
			bfwalk(node.GetRight(), max, level, seen)
		}

		if node.GetLeft() != nil {
			bfwalk(node.GetLeft(), max, level, seen)
		}
	}
}

// BFSearch is a breadth first search for node containing `value`
func BFSearch(root NodeComponent, value int64) NodeComponent {
	/*

				tree
				----
				  5         <-- level 0
				/   \
			  3      8      <-- level 1
			 / \    /  \
		   1    4  6    9   <-- level 2
		  / \
			 2              <-- level 3

	*/
	// Level 0
	if root != nil && root.GetValue() == value {
		return root
	}

	level := 1
	for {
		seen := []NodeComponent{}
		bfwalk(root, level, 0, &seen)

		for _, item := range seen {
			if item.GetValue() == value {
				return item
			}

			// This is necessary so that we do not loop forever when
			// we have already iterated over all the nodes in the tree.
			if level >= len(seen) {
				return nil
			}

		}
		level++
	}
}

// Tree represents a binary tree of nodes.
type Tree struct {
	Root NodeComponent
}

// Add adds a new leaf node to the tree.
func (t *Tree) Add(node NodeComponent) error {
	n := t.Root
	for n != node {
		// smaller then root (right)
		if node.GetValue() < n.GetValue() {
			if n.GetRight() == nil {
				n.SetRight(node)
				continue
			}

			n = n.GetRight()
			continue
		}

		// larger then root (left)
		if node.GetValue() > n.GetValue() {
			if n.GetLeft() == nil {
				n.SetLeft(node)
				continue
			}
			n = n.GetLeft()
			continue
		}

		n = nil
	}
	return nil
}

// Contains returns true if a node with `value` is found in the tree.
func (t Tree) Contains(value int64) bool {
	return BFSearch(t.Root, value) != nil
}

// NewTree returns a new tree populated with one or
// node components.
func NewTree(root NodeComponent, nodes ...NodeComponent) (*Tree, error) {
	t := Tree{Root: root}
	for _, node := range nodes {
		err := t.Add(node)
		if err != nil {
			return nil, err
		}
	}
	return &t, nil
}
