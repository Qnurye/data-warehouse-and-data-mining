package base

import "github.com/deckarep/golang-set/v2"

type Pattern struct {
	Set mapset.Set[string]
}

func NewPattern(vals ...string) Pattern {
	return Pattern{mapset.NewSet[string](vals...)}
}

func (p Pattern) String() string {
	return p.Set.String()
}

func (p Pattern) Cardinality() int {
	return p.Set.Cardinality()
}

func (p Pattern) Intersect(p2 Pattern) Pattern {
	return Pattern{p.Set.Intersect(p2.Set)}
}

func (p Pattern) IsEmpty() bool {
	return p.Set.Cardinality() == 0
}

func (p Pattern) Iter() <-chan string {
	return p.Set.Iter()
}

func (p Pattern) Clone() Pattern {
	return Pattern{p.Set.Clone()}
}

func (p Pattern) Remove(i string) {
	p.Set.Remove(i)
}

func (p Pattern) Union(p2 Pattern) Pattern {
	return Pattern{p.Set.Union(p2.Set)}
}

func (p Pattern) IsSubset(t Transaction) bool {
	for i := range p.Iter() {
		if !t.Contains(i) {
			return false
		}
	}
	return true
}
