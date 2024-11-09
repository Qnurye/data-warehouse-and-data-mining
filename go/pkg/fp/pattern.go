package fp

import (
	"strings"
)

func joinPattern(pattern []string) string {
	return strings.Join(pattern, ",")
}

func MinePatterns(tree *Tree, headerTable map[string]*Node, minSupport int) map[string]int {
	patterns := make(map[string]int)
	for item, node := range headerTable {
		support := 0
		var conditionalPatterns [][]string

		for current := node; current != nil; current = current.next {
			support += current.count
			var path []string
			for parent := current.parent; parent != nil && parent.item != ""; parent = parent.parent {
				path = append([]string{parent.item}, path...)
			}
			for i := 0; i < current.count; i++ {
				conditionalPatterns = append(conditionalPatterns, path)
			}
		}

		if support >= minSupport {
			patterns[joinPattern([]string{item})] = support
			conditionalTree, _ := BuildTree(conditionalPatterns, minSupport)
			subPatterns := MinePatterns(conditionalTree, conditionalTree.headerTable, minSupport)
			for subPattern, subSupport := range subPatterns {
				combinedPattern := append([]string{item}, strings.Split(subPattern, ",")...)
				patterns[joinPattern(combinedPattern)] = subSupport
			}
		}
	}
	return patterns
}
