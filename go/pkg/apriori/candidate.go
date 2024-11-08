package apriori

import (
	"data-mining/pkg/base"
)

func genL1(T []base.Transaction, s float64) base.PatternsWithSupport {
	var l1 []base.PatternWithSupport
	itemCount := make(map[string]int)
	length := float64(len(T))
	for _, t := range T {
		for i := range t.Iter() {
			itemCount[i]++
		}
	}
	for i, c := range itemCount {
		if float64(c)/length >= s {
			l1 = append(l1, base.PatternWithSupport{
				Pattern: base.NewPattern(i),
				Support: base.Support(float64(c) / float64(len(T))),
			})
		}
	}
	return l1
}

func generate(fp base.Patterns) base.Patterns {
	candidates := base.NewPatterns()
	fps := fp.ToSlice()
	for i := 0; i < len(fps); i++ {
		for j := i + 1; j < len(fps); j++ {
			p1 := fps[i]
			p2 := fps[j]
			if CanMerge(p1, p2) {
				c := p1.Union(p2)
				subsets := GenSubsets(c)
				if subsets.IsSubset(fp) {
					candidates.Add(c)
				}
			}
		}
	}
	return candidates
}
