package fp

import (
	"testing"
)

// TestBuildTreeEmpty tests BuildTree with empty transactions.
func TestBuildTreeEmpty(t *testing.T) {
	var transactions [][]string
	minSupport := 1
	tree, headerTable := BuildTree(transactions, minSupport)
	if tree == nil || tree.root == nil {
		t.Errorf("Expected tree root to be not nil")
	}
	if len(headerTable) != 0 {
		t.Errorf("Expected header table to be empty, got %d items", len(headerTable))
	}
}

// TestBuildTreeNoItemsMeetingMinSupport tests BuildTree when no item meets minSupport.
func TestBuildTreeNoItemsMeetingMinSupport(t *testing.T) {
	transactions := [][]string{
		{"a", "b"},
		{"c", "d"},
	}
	minSupport := 3
	_, headerTable := BuildTree(transactions, minSupport)
	if len(headerTable) != 0 {
		t.Errorf("Expected header table to be empty, got %d items", len(headerTable))
	}
}

// TestBuildTreeWithItemsMeetingMinSupport tests BuildTree with items meeting minSupport.
func TestBuildTreeWithItemsMeetingMinSupport(t *testing.T) {
	transactions := [][]string{
		{"a", "b"},
		{"a", "c"},
		{"a", "d"},
	}
	minSupport := 2
	_, headerTable := BuildTree(transactions, minSupport)
	if len(headerTable) != 1 {
		t.Errorf("Expected header table to have 1 item, got %d", len(headerTable))
	}
	if _, ok := headerTable["a"]; !ok {
		t.Errorf("Expected header table to contain item 'a'")
	}
}

// TestBuildTreeCorrectness tests BuildTree correctness with known transactions.
func TestBuildTreeCorrectness(t *testing.T) {
	transactions := [][]string{
		{"a", "b", "c"},
		{"a", "b"},
		{"a", "c"},
		{"b", "c"},
	}
	minSupport := 2
	tree, headerTable := BuildTree(transactions, minSupport)

	expectedItems := []string{"a", "b", "c"}

	if len(headerTable) != len(expectedItems) {
		t.Errorf("Expected header table to have %d items, got %d", len(expectedItems), len(headerTable))
	}

	for _, item := range expectedItems {
		if _, ok := headerTable[item]; !ok {
			t.Errorf("Expected header table to contain item '%s'", item)
		}
	}

	// Verify the structure of the tree (this is a simplified check).
	if tree.root == nil {
		t.Errorf("Expected tree root to be not nil")
	}
	if len(tree.root.children) == 0 {
		t.Errorf("Expected tree root to have children")
	}
}

// TestBuildTreeFilteredTransactions tests if transactions are correctly filtered based on minSupport.
func TestBuildTreeFilteredTransactions(t *testing.T) {
	transactions := [][]string{
		{"a", "b", "c"},
		{"b", "c"},
		{"c"},
		{"d"},
	}
	minSupport := 2
	_, headerTable := BuildTree(transactions, minSupport)

	expectedItems := []string{"b", "c"}

	if len(headerTable) != len(expectedItems) {
		t.Errorf("Expected header table to have %d items, got %d", len(expectedItems), len(headerTable))
	}

	for _, item := range expectedItems {
		if _, ok := headerTable[item]; !ok {
			t.Errorf("Expected header table to contain item '%s'", item)
		}
	}
}
