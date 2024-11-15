package main

import (
	"data-mining/internal/data"
	"data-mining/pkg/fp"
	"flag"
	"fmt"
	"log"
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
		log.Printf("Min count: %d\n", *minCount)
	}

	_, headerTable := fp.BuildTree(transactions, *minCount)
	frequentPatterns := fp.Mine(headerTable, *minCount)

	log.Println("Frequent Patterns:")
	for pattern, support := range frequentPatterns {
		fmt.Printf("{%v}: %d\n", pattern, support)
	}

	log.Printf("Execution time: %v\n", time.Since(start))
	log.Printf("Number of FP: %d\n", len(frequentPatterns))
}
