package fp

type Node struct {
	item     string
	count    int
	parent   *Node
	children map[string]*Node
	next     *Node
}

func NewFPNode(item string, parent *Node) *Node {
	return &Node{
		item:     item,
		count:    1,
		parent:   parent,
		children: make(map[string]*Node),
	}
}

type Tree struct {
	root        *Node
	headerTable map[string]*Node
}

func NewTree() *Tree {
	return &Tree{
		root:        NewFPNode("", nil),
		headerTable: make(map[string]*Node),
	}
}

func (tree *Tree) InsertTransaction(transaction []string) {
	currentNode := tree.root
	for _, item := range transaction {
		if child, exists := currentNode.children[item]; exists {
			child.count++
			currentNode = child
		} else {
			newNode := NewFPNode(item, currentNode)
			currentNode.children[item] = newNode
			currentNode = newNode

			if head, exists := tree.headerTable[item]; exists {
				for head.next != nil {
					head = head.next
				}
				head.next = newNode
			} else {
				tree.headerTable[item] = newNode
			}
		}
	}
}
