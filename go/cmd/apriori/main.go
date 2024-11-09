package main

import (
	"data-mining/internal/data"
	"data-mining/pkg/apriori"
	"data-mining/pkg/base"
	"flag"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	path := flag.String("p", "retail.dat", "path to the transactions file")
	minSupport := flag.Float64("s", 0.1, "minimum support")
	flag.Parse()

	transactions, err := data.LoadTransactions(*path)
	if err != nil {
		panic(err)
	}

	result := apriori.Run(transactions, base.Support(*minSupport))

	for p, s := range result {
		fmt.Printf("%v: %v\n", p, s)
	}

	fmt.Printf("Execution time: %v\n", time.Since(start))
}
