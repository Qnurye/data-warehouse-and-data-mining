package fp

import (
	"reflect"
	"testing"
)

func TestMinePatterns(t *testing.T) {
	tests := []struct {
		name             string
		transactions     [][]string
		minSupport       int
		expectedPatterns map[string]int
	}{
		{
			name: "Simple Case",
			transactions: [][]string{
				{"a", "b"},
				{"b", "c", "d"},
				{"a", "c", "d", "e"},
				{"a", "d", "e"},
				{"a", "b", "c"},
			},
			minSupport: 2,
			expectedPatterns: map[string]int{
				"a":     4,
				"b":     3,
				"b,a":   2,
				"c":     3,
				"c,a":   2,
				"c,b":   2,
				"d":     3,
				"d,a":   2,
				"d,c":   2,
				"e":     2,
				"e,a":   2,
				"e,d":   2,
				"e,d,a": 2,
			},
		},
		{
			name: "Single Transaction",
			transactions: [][]string{
				{"a", "b", "c"},
			},
			minSupport: 1,
			expectedPatterns: map[string]int{
				"a":     1,
				"b":     1,
				"c":     1,
				"b,a":   1,
				"c,a":   1,
				"c,b":   1,
				"c,b,a": 1,
			},
		},
		{
			name: "No Frequent Items",
			transactions: [][]string{
				{"a"},
				{"b"},
				{"c"},
			},
			minSupport:       2,
			expectedPatterns: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree, _ := BuildTree(tt.transactions, tt.minSupport)
			patterns := MinePatterns(tree, tree.headerTable, tt.minSupport)
			if !reflect.DeepEqual(patterns, tt.expectedPatterns) {
				t.Errorf("Expected patterns %v, got %v", tt.expectedPatterns, patterns)
			}
		})
	}
}
