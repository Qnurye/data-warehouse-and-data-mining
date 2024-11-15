package apriori

import "sort"

type Transaction struct {
	items []string
	next  *Transaction
}

func (t *Transaction) isNil() bool {
	return t.items == nil && t.next == nil
}

// sortItems 依据字典序对事务进行排序
func sortItems(items []string) []string {
	sortedItems := make([]string, len(items))
	copy(sortedItems, items)
	sort.Strings(sortedItems)
	return sortedItems
}

// BuildTransactions 构建事务链表
func BuildTransactions(transactions [][]string) (*Transaction, int) {
	root := &Transaction{}
	current := root

	for _, items := range transactions {
		//if len(items) == 0 {
		//	continue
		//}

		current.items = sortItems(items)
		current.next = &Transaction{}
		current = current.next
	}

	return root, len(transactions)
}
