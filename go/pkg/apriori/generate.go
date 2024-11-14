package apriori

import "sort"

// genL1 generates the single itemsets from the transactions.
func genL1(tHead Transaction, s int) map[string]int {
	l1 := make(map[string]int)
	for !tHead.isNil() {
		for _, item := range tHead.items {
			l1[join([]string{item})]++
		}
		tHead = *tHead.next
	}

	for item, count := range l1 {
		if count < s {
			delete(l1, item)
		}
	}

	return l1
}

// genL generates the frequent itemsets from the transactions.
func genL(c [][]string, tHead Transaction, s int) map[string]int {
	l := make(map[string]int)
	for !tHead.isNil() {
		for _, pattern := range c {
			if contains(tHead.items, pattern) {
				l[join(pattern)]++
			}
		}
		tHead = *tHead.next
	}

	for item, count := range l {
		if count < s {
			delete(l, item)
		}
	}

	return l
}

// genSubPatterns generates the sub-patterns from the conditional tree.
func genSubPatterns(p []string) [][]string {
	var subPatterns [][]string
	for i := 0; i < len(p); i++ {
		subPattern := make([]string, 0, len(p)-1)
		for j := 0; j < len(p); j++ {
			if i != j {
				subPattern = append(subPattern, p[j])
			}
		}
		subPatterns = append(subPatterns, subPattern)
	}
	return subPatterns
}

// genC generates the candidate itemsets from the frequent itemsets.
func genC(l map[string]int) [][]string {
	var c [][]string
	mapList := make([]string, 0, len(l))
	for item := range l {
		mapList = append(mapList, item)
	}
	// sort mapList
	sort.Strings(mapList)
	size := len(mapList)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			p1 := split(mapList[i])
			p2 := split(mapList[j])
			if len(p1) == 1 {
				c = append(c, []string{p1[0], p2[0]})
			} else {
				if canMerge(p1, p2) {
					combined := append(p1, p2[len(p2)-1])
					sub := genSubPatterns(combined)
					valid := true
					for _, s := range sub {
						if _, ok := l[join(s)]; !ok {
							valid = false
							break
						}
					}
					if valid {
						c = append(c, combined)
					}
				}
			}
		}
	}
	return c
}
