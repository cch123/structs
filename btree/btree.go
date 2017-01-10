package btree

// Node is a node of a btree
type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

const (
	// MIDORDER tranversal mid order
	MIDORDER = iota
	// PREORDER tranversal pre order
	PREORDER
	// POSTORDER tranversal post order
	POSTORDER
)

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

// Insert will insert a node to binary tree
func (tree *BTree) Insert(data int) {
}

// Delete will delete the node, the value of witch is data
func (tree *BTree) Delete(data int) {
}

// TranversalPrint will tranvesal the tree and print the node data
// the order is defined as MIDORDER/PREORDER/POSTORDER
func (tree *BTree) TranversalPrint(tranversalType int) {
}
