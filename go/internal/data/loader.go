package data

import (
	"bufio"
	"data-mining/pkg/base"
	"log"
	"os"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func LoadTransactions(filePath string) ([]base.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}(file)

	var transactions []base.Transaction
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		set := mapset.NewSet[string]()
		for _, elem := range elements {
			set.Add(elem)
		}
		transactions = append(transactions, base.Transaction{Set: set})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func LoadTransactionsAsString(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}(file)

	var transactions [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		set := mapset.NewSet[string]()
		for _, elem := range elements {
			set.Add(elem)
		}
		transactions = append(transactions, elements)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
