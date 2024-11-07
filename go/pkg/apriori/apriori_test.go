package apriori

import (
	mapset "github.com/deckarep/golang-set/v2"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		T []transaction
		s float64
	}
	tests := []struct {
		name string
		args args
		want []patternWithSupport
	}{
		{"test", args{
			[]transaction{
				mapset.NewSet("a", "b", "c"),
				mapset.NewSet("a", "b"),
			},
			0.6,
		}, []patternWithSupport{
			{mapset.NewSet("a"), 1},
			{mapset.NewSet("b"), 1},
			{mapset.NewSet("a", "b"), 1},
		}},
		{"test2", args{
			[]transaction{
				mapset.NewSet("a", "b", "c"),
				mapset.NewSet("a", "b"),
				mapset.NewSet("d", "e"),
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "d", "e"),
				mapset.NewSet("b", "e"),
			},
			0.5,
		}, []patternWithSupport{
			{mapset.NewSet("a"), support(4) / support(6)},
			{mapset.NewSet("b"), support(4) / support(6)},
			{mapset.NewSet("e"), support(3) / support(6)},
			{mapset.NewSet("a", "b"), support(3) / support(6)},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.T, tt.args.s); !comparePatternWithSupports(got, tt.want) {
				t.Errorf("run() = %v (%v), want %v (%v)",
					extract(got), extractSupport(got), extract(tt.want), extractSupport(tt.want))
			}
		})
	}
}
