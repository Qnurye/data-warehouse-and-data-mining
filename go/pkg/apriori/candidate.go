package apriori

import mapset "github.com/deckarep/golang-set/v2"

func genL1(T []transaction, s float64) []patternWithSupport {
	var l1 []patternWithSupport
	itemCount := make(map[string]int)
	length := float64(len(T))
	for _, t := range T {
		for i := range t.Iter() {
			itemCount[i]++
		}
	}
	for i, c := range itemCount {
		if float64(c)/length >= s {
			l1 = append(l1, patternWithSupport{
				pattern: mapset.NewSet(i),
				support: support(float64(c) / float64(len(T))),
			})
		}
	}
	return l1
}

func generate(fp patterns) patterns {
	candidates := emptyPatterns()
	fps := fp.ToSlice()
	for i := 0; i < len(fps); i++ {
		for j := i + 1; j < len(fps); j++ {
			p1 := fps[i]
			p2 := fps[j]
			if canMerge(p1, p2) {
				c := p1.Union(p2)
				if isSubPatterns(genSubsets(c), fp) {
					patternsAppend(candidates, c)
				}
			}
		}
	}
	return candidates
}

func canMerge(p1, p2 pattern) bool {
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

func genSubsets(p pattern) patterns {
	subsets := emptyPatterns()
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
