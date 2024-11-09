package base

import "github.com/deckarep/golang-set/v2"

type Transaction struct {
	Set mapset.Set[string]
}

func (t Transaction) Iter() <-chan string {
	return t.Set.Iter()
}

func NewTransaction(items ...string) Transaction {
	return Transaction{mapset.NewSet(items...)}
}

func (t Transaction) Contains(item string) bool {
	return t.Set.Contains(item)
}

func (t Transaction) Cardinality() int {
	return t.Set.Cardinality()
}
