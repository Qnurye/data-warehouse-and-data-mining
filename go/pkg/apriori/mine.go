package apriori

import "fmt"

func Mine(transaction Transaction, minSupport int) map[string]int {
	L := genL1(transaction, minSupport)
	fmt.Printf("Generating L1\n")
	result := L
	k := 2
	for len(L) > 0 {
		fmt.Printf("Generating L%d\n", k)
		C := genC(L)
		L = genL(C, transaction, minSupport)
		result = merge(result, L)
		k++
	}
	return result
}
