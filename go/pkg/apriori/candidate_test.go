package apriori

import (
	"data-mining/pkg/base"
	"testing"
)

func Benchmark_genL1(b *testing.B) {
	T := []base.Transaction{
		base.NewTransaction("a", "b", "c"),
		base.NewTransaction("a", "c", "d"),
		base.NewTransaction("b", "c", "e"),
	}
	s := base.Support(0.5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genL1(T, s)
	}
}

func Benchmark_generate(b *testing.B) {
	fp := base.NewPatterns(
		base.NewPattern("a", "b", "c", "d", "e"),
		base.NewPattern("a", "b", "c", "d", "f"),
		base.NewPattern("a", "b", "c", "d", "h"),
		base.NewPattern("a", "b", "c", "h", "i"),
		base.NewPattern("a", "b", "c", "h", "j"),
		base.NewPattern("a", "b", "c", "h", "k"),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generate(fp)
	}
}

func Test_genL1(t *testing.T) {
	type args struct {
		T []base.Transaction
		s base.Support
	}
	tests := []struct {
		name string
		args args
		want base.Patterns
	}{
		{
			"empty transactions",
			args{
				[]base.Transaction{},
				0.5,
			},
			base.NewPatterns(),
		},
		{
			"{a, b}, {a, c}, 0.6",
			args{
				[]base.Transaction{
					base.NewTransaction("a", "b"),
					base.NewTransaction("a", "c"),
				},
				0.6,
			},
			base.NewPatterns(
				base.NewPattern("a"),
			),
		},
		{
			"{a, b, c}, {d, e, f}, 0.6",
			args{
				[]base.Transaction{
					base.NewTransaction("a", "b", "c"),
					base.NewTransaction("d", "e", "f"),
				},
				0.6,
			},
			base.NewPatterns(),
		},
		{
			"{a, b, c}, {d, e, f}, {a}, 0.6",
			args{
				[]base.Transaction{
					base.NewTransaction("a", "b", "c"),
					base.NewTransaction("d", "e", "f"),
					base.NewTransaction("a"),
				},
				0.6,
			},
			base.NewPatterns(
				base.NewPattern("a"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := genL1(tt.args.T, tt.args.s)
			extract := got.Extract()
			if !extract.Equal(tt.want) {
				t.Errorf("genL1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generate(t *testing.T) {
	type args struct {
		fp base.Patterns
	}
	tests := []struct {
		name string
		args args
		want base.Patterns
	}{
		{"empty Patterns", args{
			base.NewPatterns(),
		}, base.NewPatterns()},
		{"k = 1", args{
			base.NewPatterns(
				base.NewPattern("a"),
				base.NewPattern("b"),
				base.NewPattern("c"),
			),
		}, base.NewPatterns(
			base.NewPattern("a", "b"),
			base.NewPattern("a", "c"),
			base.NewPattern("b", "c"),
		)},
		{"k = 2", args{
			base.NewPatterns(
				base.NewPattern("a", "b"),
				base.NewPattern("a", "c"),
				base.NewPattern("b", "c"),
				base.NewPattern("b", "d"),
				base.NewPattern("c", "d"),
			),
		}, base.NewPatterns(
			base.NewPattern("a", "b", "c"),
			base.NewPattern("b", "c", "d"),
		)},
		{"another k = 2", args{
			base.NewPatterns(
				base.NewPattern("a", "b"),
				base.NewPattern("a", "c"),
				base.NewPattern("b", "c"),
				base.NewPattern("b", "d"),
			),
		}, base.NewPatterns(
			base.NewPattern("a", "b", "c"),
		)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.fp); !got.Equal(tt.want) {
				t.Errorf("%s: generate() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
