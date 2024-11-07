package apriori

func run(T []transaction, s float64) []patternWithSupport {
	var r []patternWithSupport
	l := genL1(T, s)
	r = append(r, l...)
	for len(l) > 0 {
		c := generate(extract(l))
		l = genL(T, s, c)
		r = append(r, l...)
	}
	return r
}
