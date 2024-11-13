package apriori

import (
	"strings"
)

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

func joinPattern(pattern []string) string {
	return strings.Join(pattern, ",")
}

func splitPattern(pattern string) []string {
	return strings.Split(pattern, ",")
}
