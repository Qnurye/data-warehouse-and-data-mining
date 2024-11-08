package apriori

import "data-mining/pkg/base"

func count(T []base.Transaction, c base.Pattern) int {
	count := 0
	for _, t := range T {
		if c.IsSubset(t) {
			count++
		}
	}
	return count
}

func genL(T []base.Transaction, s float64, C base.Patterns) base.PatternsWithSupport {
	var l = base.NewPatternsWithSupport()
	total := float64(len(T))
	for c := range C.Iter() {
		count := float64(count(T, c))
		if count/total >= s {
			l = l.Add(base.PatternWithSupport{Pattern: c, Support: base.Support(count / total)})
		}
	}
	return l
}

func CanMerge(p1, p2 base.Pattern) bool {
	if p1.Cardinality() != p2.Cardinality() {
		return false
	}
	if p1.Cardinality() == 1 {
		if p1.Intersect(p2).Cardinality() == 0 {
			return true
		} else {
			return false
		}
	}
	return p1.Intersect(p2).Cardinality() == p1.Cardinality()-1
}

func GenSubsets(p base.Pattern) base.Patterns {
	subsets := base.NewPatterns()
	if p.IsEmpty() || p.Cardinality() == 1 {
		return subsets
	}
	for i := range p.Iter() {
		s := p.Clone()
		s.Remove(i)
		subsets.Add(s)
	}
	return subsets
}
