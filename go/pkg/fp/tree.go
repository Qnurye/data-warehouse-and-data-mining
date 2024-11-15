package fp

// Node 代表 FP 树中的一个节点
type Node struct {
	item     string           // 项目名称
	count    int              // 项目计数
	parent   *Node            // 父节点
	children map[string]*Node // 子节点
	next     *Node            // 链接到相同项目的下一个节点
}

// NewFPNode 创建一个新的 FP 树节点
func NewFPNode(item string, parent *Node) *Node {
	return &Node{
		item:     item,
		count:    1,
		parent:   parent,
		children: make(map[string]*Node),
	}
}

// Tree 代表 FP 树
type Tree struct {
	root        *Node            // 树的根节点
	headerTable map[string]*Node // 头表，用于存储每个项目的链表头
}

// NewTree 创建一个新的 FP 树
func NewTree() *Tree {
	return &Tree{
		root:        NewFPNode("", nil),
		headerTable: make(map[string]*Node),
	}
}

// Insert 插入一笔事务到 FP 树中
func (tree *Tree) Insert(transaction []string) {
	currentNode := tree.root

	for _, item := range transaction {

		if child, exists := currentNode.children[item]; exists {
			// 老路
			child.count++
			currentNode = child
		} else {
			// 新路
			newNode := NewFPNode(item, currentNode)
			currentNode.children[item] = newNode
			currentNode = newNode

			if head, exists := tree.headerTable[item]; exists {
				// 更新头表
				for head.next != nil {
					head = head.next
				}
				head.next = newNode
			} else {
				// 新建头表
				tree.headerTable[item] = newNode
			}
		}

	}
}

// BuildTree 根据给定的事务和最小支持度构建 FP 树
func BuildTree(transactions [][]string, minSupport int) (*Tree, map[string]*Node) {

	// 扫描事务，统计每个项目的计数
	countMap := make(map[string]int)
	for _, transaction := range transactions {
		for _, item := range transaction {
			countMap[item]++
		}
	}

	// 过滤掉支持度小于最小支持度的项目
	var filtered [][]string
	for _, t := range transactions {
		var filteredT []string
		for _, i := range t {
			if countMap[i] >= minSupport {
				filteredT = append(filteredT, i)
			}
		}
		filtered = append(filtered, filteredT)
	}

	// 构建 FP 树
	fpTree := NewTree()
	for _, transaction := range filtered {
		fpTree.Insert(transaction)
	}

	return fpTree, fpTree.headerTable
}
