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
