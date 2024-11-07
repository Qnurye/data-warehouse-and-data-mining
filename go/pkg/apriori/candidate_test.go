package apriori

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
)

func Test_genL1(t *testing.T) {
	type args struct {
		T []transaction
		s float64
	}
	tests := []struct {
		name string
		args args
		want patterns
	}{
		{
			"empty transactions",
			args{
				[]transaction{},
				0.5,
			},
			emptyPatterns(),
		},
		{
			"{a, b}, {a, c}, 0.6",
			args{
				[]transaction{
					mapset.NewSet("a", "b"),
					mapset.NewSet("a", "c"),
				},
				0.6,
			},
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
		},
		{
			"{a, b, c}, {d, e, f}, 0.6",
			args{
				[]transaction{
					mapset.NewSet("a", "b", "c"),
					mapset.NewSet("d", "e", "f"),
				},
				0.6,
			},
			emptyPatterns(),
		},
		{
			"{a, b, c}, {d, e, f}, {a}, 0.6",
			args{
				[]transaction{
					mapset.NewSet("a", "b", "c"),
					mapset.NewSet("d", "e", "f"),
					mapset.NewSet("a"),
				},
				0.6,
			},
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genL1(tt.args.T, tt.args.s); !patternsEqual(extract(got), tt.want) {
				t.Errorf("genL1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canMerge(t *testing.T) {
	type args struct {
		p1 pattern
		p2 pattern
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"two patterns with different cardinality",
			args{
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "b", "c"),
			},
			false,
		}, {
			"two patterns with the same cardinality",
			args{
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "c"),
			},
			true,
		}, {
			"two patterns with the same cardinality but no common items",
			args{
				mapset.NewSet("a", "b"),
				mapset.NewSet("c", "d"),
			},
			false,
		}, {
			"identical patterns",
			args{
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "b"),
			},
			false,
		}, {
			"single item patterns",
			args{
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMerge(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("canMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genSubsets(t *testing.T) {
	type args struct {
		p pattern
	}
	tests := []struct {
		name string
		args args
		want patterns
	}{
		{
			"empty pattern",
			args{
				emptyPattern(),
			},
			emptyPatterns(),
		},
		{
			"pattern with one item",
			args{
				mapset.NewSet("a"),
			},
			emptyPatterns(),
		},
		{
			"pattern with two items",
			args{
				mapset.NewSet("a", "b"),
			},
			mapset.NewSet[pattern](
				mapset.NewSet("b"),
				mapset.NewSet("a"),
			),
		},
		{
			"pattern with three items",
			args{
				mapset.NewSet("a", "b", "c"),
			},
			mapset.NewSet[pattern](
				mapset.NewSet("a", "c"),
				mapset.NewSet("a", "b"),
				mapset.NewSet("c", "b"),
			),
		},
		{
			"pattern with four items",
			args{
				mapset.NewSet("a", "b", "c", "d"),
			},
			mapset.NewSet[pattern](
				mapset.NewSet("a", "c", "d"),
				mapset.NewSet("a", "b", "d"),
				mapset.NewSet("a", "b", "c"),
				mapset.NewSet("c", "b", "d"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genSubsets(tt.args.p); !patternsEqual(got, tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generate(t *testing.T) {
	type args struct {
		fp patterns
	}
	tests := []struct {
		name string
		args args
		want patterns
	}{
		{"empty patterns", args{
			emptyPatterns(),
		}, emptyPatterns()},
		{"k = 1", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
				mapset.NewSet("c"),
			),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a", "b"),
			mapset.NewSet("a", "c"),
			mapset.NewSet("b", "c"),
		)},
		{"k = 2", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "c"),
				mapset.NewSet("b", "c"),
				mapset.NewSet("b", "d"),
				mapset.NewSet("c", "d"),
			),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a", "b", "c"),
			mapset.NewSet("b", "c", "d"),
		)},
		{"another k = 2", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a", "b"),
				mapset.NewSet("a", "c"),
				mapset.NewSet("b", "c"),
				mapset.NewSet("b", "d"),
			),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a", "b", "c"),
		)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.fp); !patternsEqual(got, tt.want) {
				t.Errorf("%s: generate() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
