package base

type Support float64

type PatternWithSupport struct {
	Pattern Pattern
	Support Support
}

type PatternsWithSupport map[Pattern]Support

func (pws PatternWithSupport) Equal(qws PatternWithSupport) bool {
	return pws.Pattern.Set.Equal(qws.Pattern.Set) && pws.Support == qws.Support
}

func NewPatternsWithSupport() PatternsWithSupport {
	return make(PatternsWithSupport)
}

func WithSupport(p Patterns) PatternsWithSupport {
	pws := NewPatternsWithSupport()
	for _, pattern := range p.ToSlice() {
		pws[pattern] = 0
	}
	return pws
}

func (pws PatternsWithSupport) Append(p PatternWithSupport) PatternsWithSupport {
	pws[p.Pattern] = p.Support
	return pws
}

func (pws PatternsWithSupport) Add(qws PatternsWithSupport) PatternsWithSupport {
	for k, v := range qws {
		if _, ok := pws[k]; !ok {
			pws[k] = v
		} else {
			pws[k] += v
		}
	}
	return pws
}

func (pws PatternsWithSupport) Equal(qws PatternsWithSupport) bool {
	if len(pws) != len(qws) {
		return false
	}
	for k, v := range pws {
		for qk, qv := range qws {
			if k.Equal(qk) {
				if v == qv {
					break
				}
				return false
			}
		}
	}
	return true
}

func (pws PatternsWithSupport) ExtractPatterns() Patterns {
	p := NewPatterns()
	for pattern := range pws {
		p.Add(pattern)
	}
	return p
}

func (pws PatternsWithSupport) Extract() Patterns {
	return pws.ExtractPatterns()
}

func (pws PatternsWithSupport) ExtractSupport() []Support {
	var p []Support
	for _, s := range pws {
		p = append(p, s)
	}
	return p
}
