package trees

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{Value: value, Left: nil, Right: nil}
}

func (n *BinaryTreeNode) Insert(value int) {
	if value > n.Value {
		if n.Right != nil {
			n.Right.Insert(value)
		} else {
			n.Right = NewNode(value)
		}
	} else {
		if n.Left != nil {
			n.Left.Insert(value)
		} else {
			n.Left = NewNode(value)
		}
	}
}

func (n *BinaryTreeNode) Search(value int) *BinaryTreeNode {
	if n.Value == value {
		return n
	}

	if value > n.Value {
		if n.Right == nil {
			return nil
		}

		return n.Right.Search(value)
	} else {
		if n.Left == nil {
			return nil
		}

		return n.Left.Search(value)
	}
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() BinaryTree {
	return BinaryTree{Root: nil}
}

func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = NewNode(value)
		return
	}
	t.Root.Insert(value)
}

func (t BinaryTree) Search(value int) *BinaryTreeNode {
	if t.Root == nil {
		return nil
	}

	return t.Root.Search(value)
}
