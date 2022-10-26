package trees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value, Left: nil, Right: nil}
}

func (n *Node) Insert(value int) {
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

func (n *Node) Search(value int) *Node {
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
	Root *Node
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

func (t BinaryTree) Search(value int) *Node {
	if t.Root == nil {
		return nil
	}

	return t.Root.Search(value)
}
