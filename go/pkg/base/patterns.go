package base

import "github.com/deckarep/golang-set/v2"

type Patterns struct {
	Set mapset.Set[Pattern]
}

func (ps Patterns) String() string {
	var s = "["
	for p := range ps.Set.Iter() {
		s += p.String() + ", "
	}
	if len(s) > 1 {
		s = s[:len(s)-2]
	}
	s += "]"
	return s
}

func (ps Patterns) Cardinality() int {
	return ps.Set.Cardinality()
}

func NewPatterns(p ...Pattern) Patterns {
	s := mapset.NewSet[Pattern]()
	for _, i := range p {
		s.Add(i)
	}
	return Patterns{s}
}

func (ps Patterns) ToSlice() []Pattern {
	var p []Pattern
	for i := range ps.Set.Iter() {
		p = append(p, i)
	}
	return p
}

func (ps Patterns) Add(q Pattern) bool {
	if !(ps.Contains(q)) {
		return ps.Set.Add(q)
	}
	return false
}

func (ps Patterns) Iter() <-chan Pattern {
	return ps.Set.Iter()
}

func (ps Patterns) Contains(q Pattern) bool {
	psl := ps.Set.ToSlice()
	for _, r := range psl {
		if r.Set.Equal(q.Set) {
			return true
		}
	}
	return false
}

func (ps Patterns) IsSubset(qs Patterns) bool {
	if ps.Set.Cardinality() > qs.Set.Cardinality() {
		return false
	}
	psl := ps.ToSlice()
	for _, p := range psl {
		if !qs.Contains(p) {
			return false
		}
	}
	return true
}

func (ps Patterns) Equal(qs Patterns) bool {
	if ps.Set.Cardinality() != qs.Set.Cardinality() {
		return false
	}
	return ps.IsSubset(qs)
}
