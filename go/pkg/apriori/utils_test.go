package apriori

import (
	"data-mining/pkg/base"
	"testing"
)

func Benchmark_genL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		genL([]base.Transaction{
			base.NewTransaction("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewTransaction("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewTransaction("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewTransaction("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewTransaction("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewTransaction("a", "b"),
			base.NewTransaction("a", "b", "d"),
			base.NewTransaction("a", "b", "c"),
			base.NewTransaction("a", "b", "c"),
			base.NewTransaction("a", "b", "c"),
		}, 0.1, base.NewPatterns(
			base.NewPattern("a"),
			base.NewPattern("b"),
			base.NewPattern("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
		))
	}
}

func Benchmark_genSubsets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenSubsets(base.NewPattern("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"))
	}
}

func Benchmark_canMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CanMerge(base.NewPattern("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"),
			base.NewPattern("a", "b", "c", "d", "e", "f", "g", "h", "i", "j"))
	}
}

func Test_canMerge(t *testing.T) {
	type args struct {
		p1 base.Pattern
		p2 base.Pattern
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"two Patterns with different cardinality",
			args{
				base.NewPattern("a", "b"),
				base.NewPattern("a", "b", "c"),
			},
			false,
		}, {
			"two Patterns with the same cardinality",
			args{
				base.NewPattern("a", "b"),
				base.NewPattern("a", "c"),
			},
			true,
		}, {
			"two Patterns with the same cardinality but no common items",
			args{
				base.NewPattern("a", "b"),
				base.NewPattern("c", "d"),
			},
			false,
		}, {
			"identical Patterns",
			args{
				base.NewPattern("a", "b"),
				base.NewPattern("a", "b"),
			},
			false,
		}, {
			"single item Patterns",
			args{
				base.NewPattern("a"),
				base.NewPattern("b"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanMerge(tt.args.p1, tt.args.p2); got != tt.want {
				t.Errorf("CanMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genSubsets(t *testing.T) {
	type args struct {
		p base.Pattern
	}
	tests := []struct {
		name string
		args args
		want base.Patterns
	}{
		{
			"empty Pattern",
			args{
				base.NewPattern(),
			},
			base.NewPatterns(),
		},
		{
			"Pattern with one item",
			args{
				base.NewPattern("a"),
			},
			base.NewPatterns(),
		},
		{
			"Pattern with two items",
			args{
				base.NewPattern("a", "b"),
			},
			base.NewPatterns(
				base.NewPattern("b"),
				base.NewPattern("a"),
			),
		},
		{
			"Pattern with three items",
			args{
				base.NewPattern("a", "b", "c"),
			},
			base.NewPatterns(
				base.NewPattern("a", "c"),
				base.NewPattern("a", "b"),
				base.NewPattern("c", "b"),
			),
		},
		{
			"Pattern with four items",
			args{
				base.NewPattern("a", "b", "c", "d"),
			},
			base.NewPatterns(
				base.NewPattern("a", "c", "d"),
				base.NewPattern("a", "b", "d"),
				base.NewPattern("a", "b", "c"),
				base.NewPattern("c", "b", "d"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenSubsets(tt.args.p); !got.Equal(tt.want) {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
