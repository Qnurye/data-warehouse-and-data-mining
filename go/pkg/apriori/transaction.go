package apriori

import "sort"

type Transaction struct {
	items []string
	next  *Transaction
}

func (t *Transaction) Items() []string {
	return t.items
}

func (t *Transaction) Next() *Transaction {
	return t.next
}

func (t *Transaction) isNil() bool {
	return t.items == nil && t.next == nil
}

// sortItems sorts the items in the transaction in lexicographical order.
func sortItems(items []string) []string {
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)
	sort.Strings(sortedItems)
	return sortedItems
}

// BuildTransactions builds a linked list of transactions from a list of transactions.
func BuildTransactions(transactions [][]string) (*Transaction, int) {
	root := &Transaction{}
	current := root

	for _, items := range transactions {
		if len(items) == 0 {
			continue
		}

		current.items = sortItems(items)
		current.next = &Transaction{}
		current = current.next
	}

	return root, len(transactions)
}
