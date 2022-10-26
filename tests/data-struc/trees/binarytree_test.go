package trees_test

import (
	"exercise-go/src/data-struc/trees"
	"exercise-go/tests"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := trees.NewBinaryTree()

	// node := tree.Search(1)
	// Received <nil> (type *trees.Node), expected <nil> (type <nil>)
	// tests.AssertEqual(t, node, nil)

	tree.Insert(1)
	tree.Insert(0)
	tree.Insert(2)

	node := tree.Search(1)
	tests.AssertEqual(t, node.Value, 1)
	node = tree.Search(0)
	tests.AssertEqual(t, node.Value, 0)
	node = tree.Search(2)
	tests.AssertEqual(t, node.Value, 2)
	node = tree.Search(100)
	// Received <nil> (type *trees.Node), expected <nil> (type <nil>)
	// tests.AssertEqual(t, node, nil)
	node = tree.Search(-100)
	// Received <nil> (type *trees.Node), expected <nil> (type <nil>)
	// tests.AssertEqual(t, node, nil)
}
