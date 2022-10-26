package trees

type TrieNode struct {
	value    byte
	children []*TrieNode
	end      bool
}

func NewTrieNode(value string) *TrieNode {
	if value == "" {
		return &TrieNode{value: 0, children: []*TrieNode{}, end: false}
	}

	var node *TrieNode
	end := len(value) == 1

	node = &TrieNode{value: value[0], children: []*TrieNode{}, end: end}
	node.Insert(value[1:])

	return node
}

func (n *TrieNode) contains(char string) *TrieNode {
	// TODO: Replace with efficient search
	for _, v := range n.children {
		if v.value == char[0] {
			return v
		}
	}

	return nil
}

func (n *TrieNode) Insert(value string) {
	if value == "" {
		return
	}

	if v := n.contains(value[:1]); v != nil {
		v.Insert(value[1:])
	}

	n.children = append(n.children, NewTrieNode(value))
}

func (n *TrieNode) Search(value string) bool {
	if value == "" {
		return false
	}

	if v := n.contains(value[:1]); v != nil {
		if len(value) == 1 && v.end {
			return true
		}

		return v.Search(value[1:])
	}

	return false
}

type Trie struct {
	root *TrieNode
}

func NewTrie() Trie {
	return Trie{root: NewTrieNode("")}
}

func (t *Trie) Insert(value string) {
	t.root.Insert(value)
}

func (t *Trie) Search(value string) bool {
	// TODO: write an actual useful search that returns list of strings for a given prefix
	return t.root.Search(value)
}
