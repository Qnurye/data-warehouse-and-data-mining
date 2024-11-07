package apriori

import (
	mapset "github.com/deckarep/golang-set/v2"
	"testing"
)

func Test_patternsContain(t *testing.T) {
	type args struct {
		a patterns
		b pattern
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty patterns", args{
			emptyPatterns(),
			emptyPattern(),
		}, false},
		{"a in {a}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
			mapset.NewSet("a"),
		}, true},
		{"ab in {a, b}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet("a", "b"),
		}, false},
		{"a in {a, b, c}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
				mapset.NewSet("c"),
			),
			mapset.NewSet("a"),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patternsContain(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("patternsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_patternsAppend(t *testing.T) {
	type args struct {
		a patterns
		b pattern
	}
	tests := []struct {
		name string
		args args
		want patterns
	}{
		{"empty patterns", args{
			emptyPatterns(),
			emptyPattern(),
		}, mapset.NewSet[pattern](emptyPattern())},
		{"a to {a}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
			mapset.NewSet("a"),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a"),
		)},
		{"ab to {a, b}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet("a", "b"),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a"),
			mapset.NewSet("b"),
			mapset.NewSet("a", "b"),
		)},
		{"a to {a, b, c}", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
				mapset.NewSet("c"),
			),
			mapset.NewSet("a"),
		}, mapset.NewSet[pattern](
			mapset.NewSet("a"),
			mapset.NewSet("b"),
			mapset.NewSet("c"),
		)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _ = patternsAppend(tt.args.a, tt.args.b); !patternsEqual(tt.args.a, tt.want) {
				t.Errorf("patternsEqual() = %v, want %v", tt.args.a, tt.want)
			}
		})
	}
}

func Test_isSubPatterns(t *testing.T) {
	type args struct {
		a patterns
		b patterns
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty patterns", args{
			emptyPatterns(),
			emptyPatterns(),
		}, true},
		{"patterns with different cardinality", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("a", "b"),
			),
		}, true},
		{"patterns with the same cardinality but different items", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("c"),
			),
		}, false},
		{"identical patterns", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubPatterns(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("patternsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_patternsEqual(t *testing.T) {
	type args struct {
		a patterns
		b patterns
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty patterns", args{
			emptyPatterns(),
			emptyPatterns(),
		}, true},
		{"patterns with different cardinality", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("a", "b"),
			),
		}, false},
		{"patterns with the same cardinality but different items", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("c"),
			),
		}, false},
		{"identical patterns", args{
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
			mapset.NewSet[pattern](
				mapset.NewSet("a"),
				mapset.NewSet("b"),
			),
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := patternsEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("patternsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
