package fp

import (
	"testing"
)

func TestNewFPNode(t *testing.T) {
	node := NewFPNode("item1", nil)
	if node.item != "item1" {
		t.Errorf("Expected node.item to be 'item1', got '%s'", node.item)
	}
	if node.count != 1 {
		t.Errorf("Expected node.count to be 1, got %d", node.count)
	}
	if node.parent != nil {
		t.Errorf("Expected node.parent to be nil, got %v", node.parent)
	}
	if len(node.children) != 0 {
		t.Errorf("Expected node.children to be empty, got %d", len(node.children))
	}
	if node.next != nil {
		t.Errorf("Expected node.next to be nil, got %v", node.next)
	}
}

func TestNewTree(t *testing.T) {
	tree := NewTree()
	if tree.root == nil {
		t.Errorf("Expected tree.root to be initialized")
	}
	if tree.root.item != "" {
		t.Errorf("Expected tree.root.item to be empty string, got '%s'", tree.root.item)
	}
	if tree.root.count != 1 {
		t.Errorf("Expected tree.root.count to be 1, got %d", tree.root.count)
	}
	if tree.root.parent != nil {
		t.Errorf("Expected tree.root.parent to be nil, got %v", tree.root.parent)
	}
	if len(tree.root.children) != 0 {
		t.Errorf("Expected tree.root.children to be empty, got %d", len(tree.root.children))
	}
	if len(tree.headerTable) != 0 {
		t.Errorf("Expected headerTable to be empty, got %d", len(tree.headerTable))
	}
}

func TestInsertTransaction(t *testing.T) {
	tree := NewTree()
	transaction1 := []string{"item1", "item2", "item3"}
	tree.InsertTransaction(transaction1)

	node1 := tree.root.children["item1"]
	if node1 == nil || node1.item != "item1" {
		t.Errorf("Expected child 'item1' of root")
		return
	}
	if node1.count != 1 {
		t.Errorf("Expected node1.count to be 1, got %d", node1.count)
	}

	node2 := node1.children["item2"]
	if node2 == nil || node2.item != "item2" {
		t.Errorf("Expected child 'item2' of node1")
		return
	}
	if node2.count != 1 {
		t.Errorf("Expected node2.count to be 1, got %d", node2.count)
	}

	node3 := node2.children["item3"]
	if node3 == nil || node3.item != "item3" {
		t.Errorf("Expected child 'item3' of node2")
		return
	}
	if node3.count != 1 {
		t.Errorf("Expected node3.count to be 1, got %d", node3.count)
	}

	transaction2 := []string{"item1", "item4"}
	tree.InsertTransaction(transaction2)

	if node1.count != 2 {
		t.Errorf("Expected node1.count to be 2 after second transaction, got %d", node1.count)
	}

	node4 := node1.children["item4"]
	if node4 == nil || node4.item != "item4" {
		t.Errorf("Expected child 'item4' of node1")
		return
	}
	if node4.count != 1 {
		t.Errorf("Expected node4.count to be 1, got %d", node4.count)
	}

	transaction3 := []string{"item2", "item5"}
	tree.InsertTransaction(transaction3)

	node2Alt := tree.root.children["item2"]
	if node2Alt == nil || node2Alt.item != "item2" {
		t.Errorf("Expected child 'item2' of root")
		return
	}
	if node2Alt.count != 1 {
		t.Errorf("Expected node2Alt.count to be 1, got %d", node2Alt.count)
	}

	if tree.headerTable["item2"].next != node2Alt {
		t.Errorf("Expected headerTable['item2'].next to be node2Alt")
	}

	node5 := node2Alt.children["item5"]
	if node5 == nil || node5.item != "item5" {
		t.Errorf("Expected child 'item5' of node2Alt")
		return
	}
	if node5.count != 1 {
		t.Errorf("Expected node5.count to be 1, got %d", node5.count)
	}
}
