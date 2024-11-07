package apriori

func count(T []transaction, c pattern) int {
	count := 0
	for _, t := range T {
		if c.IsSubset(t) {
			count++
		}
	}
	return count
}

func genL(T []transaction, s float64, C patterns) []patternWithSupport {
	var l []patternWithSupport
	length := float64(len(T))
	for c := range C.Iter() {
		count := float64(count(T, c))
		if count/length >= s {
			l = append(l, patternWithSupport{c, support(count / length)})
		}
	}
	return l
}
