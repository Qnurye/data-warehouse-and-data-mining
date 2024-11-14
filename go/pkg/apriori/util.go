package apriori

import (
	"strings"
)

// merge merges two maps of patterns and their support counts.
func merge(a map[string]int, b map[string]int) map[string]int {
	for pattern, support := range b {
		a[pattern] = support
	}
	return a
}

// contains returns true if the alphabetically sorted transaction contains the alphabetically sorted pattern.
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

func join(pattern []string) string {
	return strings.Join(pattern, ",")
}

func split(pattern string) []string {
	return strings.Split(pattern, ",")
}

// canMerge returns true if the two patterns can be merged.
func canMerge(p1, p2 []string) bool {
	for i := 0; i < len(p1)-1; i++ {
		if p1[i] != p2[i] {
			return false
		}
	}
	return p1[len(p1)-1] < p2[len(p2)-1]
}
