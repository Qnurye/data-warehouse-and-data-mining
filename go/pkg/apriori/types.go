package apriori

import mapset "github.com/deckarep/golang-set/v2"

type support float64
type transaction mapset.Set[string] // A transaction is a set of items
type pattern mapset.Set[string]     // A pattern is a set of items
type patterns mapset.Set[pattern]   // An itemset is a set of patterns
type patternWithSupport struct {
	pattern pattern
	support support
}

func comparePatternWithSupport(a, b patternWithSupport) bool {
	return a.pattern.Equal(b.pattern) && a.support == b.support
}

func comparePatternWithSupports(a, b []patternWithSupport) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		for j := range b {
			if comparePatternWithSupport(a[i], b[j]) {
				break
			}
			if j == len(b)-1 {
				return false
			}
		}
	}
	return true
}

func extract(pws []patternWithSupport) patterns {
	p := emptyPatterns()
	for _, pw := range pws {
		p.Add(pw.pattern)
	}
	return p
}

func extractSupport(pws []patternWithSupport) []support {
	var p []support
	for _, pw := range pws {
		p = append(p, pw.support)
	}
	return p
}

func patternsContain(p patterns, q pattern) bool {
	for r := range p.Iter() {
		if r.Equal(q) {
			return true
		}
	}
	return false
}

func isSubPatterns(a, b patterns) bool {
	for p := range a.Iter() {
		if !patternsContain(b, p) {
			return false
		}
	}
	return true
}

func patternsAppend(p patterns, q pattern) bool {
	if !(patternsContain(p, q)) {
		return p.Add(q)
	}
	return false
}

func patternsEqual(a, b patterns) bool {
	if a.Cardinality() != b.Cardinality() {
		return false
	}
	return isSubPatterns(a, b)
}
func emptyPattern() pattern {
	return mapset.NewSet[string]()
}

func emptyPatterns() patterns {
	return mapset.NewSet[pattern]()
}
