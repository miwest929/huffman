package huffmantree

type TreeNode struct {
	left      *Tree
	symbol    string
	frequency int
	right     *Tree
}

func NewTreeNode(symbol string, freq int) *TreeNode {
	return &TreeNode{Symbol: symbol, Frequency: freq, Left: nil, Right: nil}
}

func (node *TreeNode) Left() *TreeNode {
	return node.left
}

func (node *TreeNode) SetLeft(left *TreeNode) {
	node.left = left
}

func (node *TreeNode) Right() *TreeNode {
	return node.right
}

func (node *TreeNode) SetRight(right *TreeNode) {
	node.right = right
}

func (node *TreeNode) Symbol() string {
	return node.symbol
}

func (node *TreeNode) Frequency() int {
	return node.frequency
}
