package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
