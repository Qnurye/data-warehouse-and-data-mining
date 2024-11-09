package apriori

import (
	"data-mining/pkg/base"
)

func Run(T []base.Transaction, s base.Support) base.PatternsWithSupport {
	r := base.NewPatternsWithSupport()
	l := genL1(T, s)
	r.Add(l)
	k := 1
	for len(l) > 0 {
		//fmt.Println("Generating L", k+1)
		c := generate(l.Extract())
		l = genL(T, s, c)
		r.Add(l)
		k++
	}
	return r
}
