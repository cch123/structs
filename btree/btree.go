package btree

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
func New() {
	root := new(Node)
	return &BTree{
		Root: root,
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
	insertPoint := tree.findInsertPoint(data)
	newNode := new(Node)
	newNode.Data = data
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
	for cur.Left != nil || cur.Right != nil {
		if cur.Left != nil && data < cur.Data {
			parent = cur
			cur = cur.Left
			continue
		}

		if cur.Right != nil && data > cur.Data {
			parent = cur
			cur = cur.Right
			continue
		}
		break
	}

	if cur.Data == data {
		if parent != cur {
			if parent.Left.Data == data {
				parent.Left = nil
			} else {
				parent.Right = nil
			}
		}
		return true
	}

	return false
}

// TranversalPrint will tranvesal the tree and print the node data
// the order is defined as MIDORDER/PREORDER/POSTORDER
func (tree *BTree) TranversalPrint(tranversalType int) {
}
