package btree

import "fmt"

const (
	// MIDORDER tranversal mid order
	MIDORDER = iota
	// PREORDER tranversal pre order
	PREORDER
	// POSTORDER tranversal post order
	POSTORDER
)

// Node is a node of a btree
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

// BTree is binary tree
type BTree struct {
	Root *Node
}

// New makes a new binary tree
func New() *BTree {
	return &BTree{
		Root: nil,
	}
}

func (tree *BTree) findInsertPoint(data int) (node *Node) {
	var cur = tree.Root
	for cur.Left != nil || cur.Right != nil {
		if cur.Left != nil && data < cur.Data {
			cur = cur.Left
			continue
		}

		if cur.Right != nil && data > cur.Data {
			cur = cur.Right
			continue
		}

		break
	}

	return cur
}

// Insert will insert a node to binary tree
func (tree *BTree) Insert(data int) bool {
	newNode := new(Node)
	newNode.Data = data

	if tree.Root == nil {
		tree.Root = newNode
		return true
	}

	// the root is not nil
	insertPoint := tree.findInsertPoint(data)
	if data < insertPoint.Data {
		insertPoint.Left = newNode
		return true
	}

	if data > insertPoint.Data {
		insertPoint.Right = newNode
		return true
	}

	return false
}

// Delete will delete the node, the value of witch is data
func (tree *BTree) Delete(data int) bool {
	var cur = tree.Root
	var parent = tree.Root
	var flag = true // true left, false right

	for cur != nil && cur.Data != data {
		if data < cur.Data {
			parent = cur
			cur = cur.Left
			flag = true
			continue
		}

		if data > cur.Data {
			parent = cur
			cur = cur.Right
			flag = false
			continue
		}
	}

	if cur != nil && cur.Data == data {
		// means this is the root
		if cur == parent {
			if parent.Left != nil {
				parent.Left.Right = tree.Root.Right
				tree.Root = parent.Left
			} else if parent.Right != nil {
				parent.Right.Left = tree.Root.Left
				tree.Root = parent.Right
			}
			return true
		}

		// not the root
		if flag == true {
			// delete left
			parent.Left = parent.Left.Left
			return true
		}

		// delete right
		parent.Right = parent.Right.Right
		return true
	}
	return false
}

// TraversalPrint will tranvesal the tree and print the node data
// the order is defined as MIDORDER/PREORDER/POSTORDER
func (node *Node) TraversalPrint(traversalType int) {
	switch traversalType {
	case MIDORDER:
		//print left
		if node.Left != nil {
			node.Left.TraversalPrint(traversalType)
		}

		//print current
		fmt.Println(node.Data)

		//print right
		if node.Right != nil {
			node.Right.TraversalPrint(traversalType)
		}
	case PREORDER:
		//print current
		fmt.Println(node.Data)
		//print left
		if node.Left != nil {
			node.Left.TraversalPrint(traversalType)
		}

		//print right
		if node.Right != nil {
			node.Right.TraversalPrint(traversalType)
		}
	case POSTORDER:
		//print left
		if node.Left != nil {
			node.Left.TraversalPrint(traversalType)
		}

		//print right
		if node.Right != nil {
			node.Right.TraversalPrint(traversalType)
		}

		//print current
		fmt.Println(node.Data)
	}
}
