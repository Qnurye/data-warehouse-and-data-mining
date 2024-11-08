package base

type Support float64

type PatternWithSupport struct {
	Pattern Pattern
	Support Support
}
type PatternsWithSupport []PatternWithSupport

func (pws PatternWithSupport) Equal(qws PatternWithSupport) bool {
	return pws.Pattern.Set.Equal(qws.Pattern.Set) && pws.Support == qws.Support
}

func NewPatternsWithSupport() PatternsWithSupport {
	return make(PatternsWithSupport, 0)
}

func (pws PatternsWithSupport) Add(p PatternWithSupport) PatternsWithSupport {
	return append(pws, p)
}

func (pws PatternsWithSupport) Equal(qws PatternsWithSupport) bool {
	if len(pws) != len(qws) {
		return false
	}
	for i := range pws {
		for j := range qws {
			if pws[i].Equal(qws[j]) {
				break
			}
			if j == len(qws)-1 {
				return false
			}
		}
	}
	return true
}

func (pws PatternsWithSupport) ExtractPatterns() Patterns {
	p := NewPatterns()
	for _, pw := range pws {
		p.Add(pw.Pattern)
	}
	return p
}

func (pws PatternsWithSupport) Extract() Patterns {
	return pws.ExtractPatterns()
}

func (pws PatternsWithSupport) ExtractSupport() []Support {
	var p []Support
	for _, pw := range pws {
		p = append(p, pw.Support)
	}
	return p
}
