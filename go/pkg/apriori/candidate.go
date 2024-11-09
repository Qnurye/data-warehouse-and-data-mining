package apriori

import (
	"data-mining/pkg/base"
	"runtime"
	"sync"
)

func genL1(T []base.Transaction, s base.Support) base.PatternsWithSupport {
	l1 := base.NewPatternsWithSupport()
	itemCount := make(map[string]int)
	length := base.Support(len(T))
	for _, t := range T {
		for i := range t.Iter() {
			itemCount[i]++
		}
	}
	for i, c := range itemCount {
		if base.Support(c)/length >= s {
			l1.Append(base.PatternWithSupport{
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
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	tasks := make(chan [2]int, numWorkers*2)
	results := make(chan base.Pattern, numWorkers*2)

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for indices := range tasks {
				i, j := indices[0], indices[1]
				p1, p2 := fps[i], fps[j]
				if CanMerge(p1, p2) {
					c := p1.Union(p2)
					if GenSubsets(c).IsSubset(fp) {
						results <- c
					}
				}
			}
		}()
	}

	go func() {
		for i := 0; i < len(fps); i++ {
			for j := i + 1; j < len(fps); j++ {
				tasks <- [2]int{i, j}
			}
		}
		close(tasks)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for c := range results {
		candidates.Add(c)
	}

	//fmt.Printf("Total candidates generated: %d\n", candidates.Cardinality())
	return candidates
}
