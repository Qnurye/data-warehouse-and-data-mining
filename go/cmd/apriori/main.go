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
	minSupport := flag.Float64("s", 0.5, "minimum support")
	flag.Parse()

	transactions, err := data.LoadTransactions(*path)
	if err != nil {
		panic(err)
	}

	result := apriori.Run(transactions, *minSupport)

	for _, pws := range result {
		fmt.Printf("%v: %v\n", pws.Pattern, pws.Support)
	}

	fmt.Printf("Execution time: %v\n", time.Since(start))
}
