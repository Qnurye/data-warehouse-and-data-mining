package main

import (
	"data-mining/internal/data"
	"data-mining/pkg/apriori"
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	path := flag.String("p", "retail.dat", "path to the transactions file")
	minSupport := flag.Float64("s", 0.1, "minimum support")
	minCount := flag.Int("c", 0, "minimum count")
	flag.Parse()

	transactions, err := data.LoadTransactionsAsString(*path)
	if err != nil {
		panic(err)
	}

	if *minCount == 0 {
		*minCount = int(float64(len(transactions)) * *minSupport)
		fmt.Printf("Min count: %d\n", *minCount)
	}

	tHead, cnt := apriori.BuildTransactions(transactions)
	fmt.Printf("Number of transactions: %d\n", cnt)
	frequentPatterns := apriori.Mine(*tHead, *minCount)

	fmt.Println("Frequent Patterns:")
	for pattern, support := range frequentPatterns {
		fmt.Printf("Pattern: %v, Support: %d\n", pattern, support)
	}

	fmt.Printf("Execution time: %v\n", time.Since(start))
	fmt.Printf("Number of FP: %d\n", len(frequentPatterns))
}
