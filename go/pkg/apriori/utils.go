package apriori

import (
	"data-mining/pkg/base"
)

func count(T []base.Transaction, c base.Patterns) base.PatternsWithSupport {
	result := base.WithSupport(c)
	for _, t := range T {
		for _, p := range c.ToSlice() {
			if p.IsSubset(t) {
				result[p]++
			}
		}
	}

	for p, sp := range result {
		result[p] = sp / base.Support(len(T))
	}
	return result
}

func genL(T []base.Transaction, s base.Support, C base.Patterns) base.PatternsWithSupport {
	var l = base.NewPatternsWithSupport()
	if C.Cardinality() == 0 {
		return l
	}

	result := count(T, C)
	for pt, sp := range result {
		if sp >= s {
			l.Append(base.PatternWithSupport{Pattern: pt, Support: sp})
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

	ps := p.ToSlice()
	for _, i := range ps {
		s := p.Clone()
		s.Remove(i)
		subsets.Add(s)
	}

	return subsets
}
