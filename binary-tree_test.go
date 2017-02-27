package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// makeTree returns a binary tree and all the nodes in
// the order as they where inserted into the tree.
func makeTree() (*Tree, map[uint64]NodeComponent) {
	/*

	         5        Level 0
	       /   \
	      3     8     Level 1
	     / \   / \
	    1   4 6   9   Level 2
	   / \
	      2           Level 3

	*/
	nodes := make(map[uint64]NodeComponent)

	one := NewNode(1)
	nodes[1] = one
	two := NewNode(2)
	nodes[2] = two
	three := NewNode(3)
	nodes[3] = three
	four := NewNode(4)
	nodes[4] = four
	five := NewNode(5)
	nodes[5] = five
	six := NewNode(6)
	nodes[6] = six
	eight := NewNode(8)
	nodes[8] = eight
	nine := NewNode(9)
	nodes[9] = nine

	tree, err := NewTree(
		five,
		three, one, four,
		two,
		eight,
		six, nine,
	)

	if err != nil {
		log.Fatal(err)
	}

	return tree, nodes
}

func TestNewNode(t *testing.T) {
	node := NewNode(1)
	assert.Implements(t, (*NodeComponent)(nil), node)
}

func TestNodeGetValue(t *testing.T) {
	root := NewNode(1)
	assert.EqualValues(t, root.GetValue(), 1)
}

func TestNodeGetLeft(t *testing.T) {
	root := NewNode(1)
	left := NewNode(2)
	root.SetLeft(left)
	assert.Equal(t, root.GetLeft(), left)
}

func TestNodeGetRight(t *testing.T) {
	root := NewNode(1)
	right := NewNode(0)
	root.SetRight(right)
	assert.Equal(t, root.GetRight(), right)
}

func TestNewBinaryTree(t *testing.T) {
	root := NewNode(5)
	tree, err := NewTree(root)
	assert.Nil(t, err)
	assert.Equal(t, root, tree.Root)
}

func TestNewBinaryTreePassingNodes(t *testing.T) {
	/*
	       5
	      /
	     3
	    / \
	   1   4
	    \
	      2
	*/
	five := NewNode(5)
	three := NewNode(3)
	one := NewNode(1)
	four := NewNode(4)
	two := NewNode(2)
	tree, err := NewTree(five, three, one, four, two)
	assert.Nil(t, err)
	assert.Equal(t, five, tree.Root)
	assert.Equal(t, three, five.GetRight())
	assert.Equal(t, one, three.GetRight())
	assert.Equal(t, four, three.GetLeft())
	assert.Equal(t, two, one.GetLeft())
}

func TestTreeAddRight(t *testing.T) {
	five := NewNode(5)
	tree, err := NewTree(five)

	/*

	     5
	    / \
	   3

	*/
	three := NewNode(3)
	err = tree.Add(three)
	assert.Nil(t, err)
	assert.Nil(t, tree.Root.GetLeft())
	assert.Equal(t, three, tree.Root.GetRight())

	/*

	       5
	      / \
	     3
	    / \
	   1

	*/
	one := NewNode(1)
	err = tree.Add(one)
	assert.Nil(t, err)
	assert.Nil(t, tree.Root.GetLeft())
	assert.Equal(t, three, tree.Root.GetRight())
	assert.Equal(t, one, three.GetRight())

	/*

	        5
	       / \
	      3
	     / \
	    1
	   / \
	      2

	*/
	two := NewNode(2)
	err = tree.Add(two)
	assert.Nil(t, err)
	assert.Nil(t, tree.Root.GetLeft())
	assert.Equal(t, three, tree.Root.GetRight())
	assert.Equal(t, one, three.GetRight())
	assert.Nil(t, one.GetRight())
	assert.Equal(t, two, one.GetLeft())

	/*

	        5
	       / \
	      3   8
	     / \
	    1
	   / \
	      2

	*/
	eight := NewNode(8)
	err = tree.Add(eight)
	assert.Nil(t, err)
	assert.Equal(t, eight, tree.Root.GetLeft())

	/*

	         5
	       /   \
	      3     8
	     / \   / \
	    1     6
	   / \
	      2

	*/
	six := NewNode(6)
	err = tree.Add(six)
	assert.Nil(t, err)
	assert.Equal(t, eight, tree.Root.GetLeft())
	assert.Equal(t, six, eight.GetRight())

	/*

	         5
	       /   \
	      3     8
	     / \   / \
	    1     6   9
	   / \
	      2

	*/
	nine := NewNode(9)
	err = tree.Add(nine)
	assert.Nil(t, err)
	assert.Equal(t, eight, tree.Root.GetLeft())
	assert.Equal(t, six, eight.GetRight())
	assert.Equal(t, nine, eight.GetLeft())

	/*

	         5
	       /   \
	      3     8
	     / \   / \
	    1   4 6   9
	   / \
	      2

	*/
	four := NewNode(4)
	err = tree.Add(four)
	assert.Nil(t, err)
	assert.Equal(t, four, three.GetLeft())
}

func TestTreeContains(t *testing.T) {
	tree, _ := makeTree()
	assert.Equal(t, true, tree.Contains(2))
}

func TestBFSearch(t *testing.T) {
	tree, nodes := makeTree()

	// Level 0
	assert.Equal(t, nodes[5], BFSearch(tree.Root, 5))

	// Level 1
	assert.Equal(t, nodes[3], BFSearch(tree.Root, 3))
	assert.Equal(t, nodes[8], BFSearch(tree.Root, 8))

	// Level 2
	assert.Equal(t, nodes[1], BFSearch(tree.Root, 1))
	assert.Equal(t, nodes[4], BFSearch(tree.Root, 4))
	assert.Equal(t, nodes[6], BFSearch(tree.Root, 6))
	assert.Equal(t, nodes[9], BFSearch(tree.Root, 9))

	// Level 3
	assert.Equal(t, nodes[2], BFSearch(tree.Root, 2))

	// Node not found in the tree
	assert.Nil(t, BFSearch(tree.Root, 10))
}

func TestBFFlatten(t *testing.T) {
	/*
	       5       level 0
	      / \
	     3         level 1
	    / \
	   1           level 2
	*/
	one := NewNode(1)
	three := NewNode(3)
	five := NewNode(5)

	tree, err := NewTree(five, three, one)
	assert.Nil(t, err)
	assert.Equal(t, []NodeComponent{five, three, one}, BFFlatten(*tree, 3))

	// See tree layout above.
	tree, nodes := makeTree()

	/*
	         5        Level 0
	       /   \
	      3     8     Level 1
	     / \   / \
	    1   4 6   9   Level 2
	   / \
	      2           Level 3
	*/
	expected := []NodeComponent{
		nodes[5],
		nodes[3], nodes[8],
		nodes[1], nodes[4],
		nodes[2],
		nodes[6], nodes[9],
	}
	assert.Equal(t, expected, BFFlatten(*tree, 3))
}
