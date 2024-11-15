package apriori

import "strings"

// merge 合并两个 map
func merge(a map[string]int, b map[string]int) map[string]int {
	for pattern, support := range b {
		a[pattern] = support
	}
	return a
}

// contains 判断 pattern 是否在 transaction 中，都依据字典序排序
func contains(transaction []string, pattern []string) bool {
	i, j := 0, 0
	for i < len(transaction) && j < len(pattern) {
		if transaction[i] == pattern[j] {
			j++
		}
		i++
	}

	return j == len(pattern)
}

// canMerge 判断两个 pattern 是否可以合并
func canMerge(p1, p2 []string) bool {
	for i := 0; i < len(p1)-1; i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	return p1[len(p1)-1] < p2[len(p2)-1]
}

func join(pattern []string) string {
	return strings.Join(pattern, ",")
}

func split(pattern string) []string {
	return strings.Split(pattern, ",")
}
