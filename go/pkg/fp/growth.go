package fp

func BuildTree(transactions [][]string, minSupport int) (*Tree, map[string]*Node) {
	itemCount := make(map[string]int)
	for _, transaction := range transactions {
		for _, item := range transaction {
			itemCount[item]++
		}
	}

	var filteredTransactions [][]string
	for _, transaction := range transactions {
		var filteredTransaction []string
		for _, item := range transaction {
			if itemCount[item] >= minSupport {
				filteredTransaction = append(filteredTransaction, item)
			}
		}
		filteredTransactions = append(filteredTransactions, filteredTransaction)
	}

	fpTree := NewTree()
	for _, transaction := range filteredTransactions {
		fpTree.InsertTransaction(transaction)
	}
	return fpTree, fpTree.headerTable
}
