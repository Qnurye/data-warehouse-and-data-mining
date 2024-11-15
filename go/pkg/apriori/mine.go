package apriori

import (
	"log"
)

func Mine(transaction Transaction, minSupport int, logging bool) map[string]int {

	L := genL1(transaction, minSupport)
	if logging {
		log.Printf("Generating L1\n")
	}
	result := L

	k := 2
	for len(L) > 0 {
		if logging {
			log.Printf("Generating C%d\n", k)
		}
		C := genC(L)

		if logging {
			log.Printf("Generating L%d\n", k)
		}
		L = genL(C, transaction, minSupport)

		result = merge(result, L)
		k++
	}
	return result

}
