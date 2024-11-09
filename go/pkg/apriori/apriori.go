package apriori

import (
	"data-mining/pkg/base"
)

func Run(T []base.Transaction, s float64) base.PatternsWithSupport {
	var r base.PatternsWithSupport
	l := genL1(T, s)
	r = append(r, l...)
	k := 1
	for len(l) > 0 {
		//fmt.Println("Generating L", k+1)
		c := generate(l.Extract())
		l = genL(T, s, c)
		r = append(r, l...)
		k++
	}
	return r
}
