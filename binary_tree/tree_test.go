package tree

import "testing"

func TestInsert(t *testing.T) {
	tree := New()
	tree.Insert(5)
	tree.Insert(4)
	tree.Insert(1)
	tree.Insert(10)
	tree.Insert(12)
	tree.Insert(9)
	tree.Root.TraversalPrint(MIDORDER)
	tree.Delete(5)
	tree.Root.TraversalPrint(MIDORDER)
}
