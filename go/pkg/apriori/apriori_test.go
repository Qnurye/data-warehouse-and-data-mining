package apriori

import (
	"data-mining/pkg/base"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run([]base.Transaction{
			base.NewTransaction("a", "b", "c", "d", "e"),
			base.NewTransaction("a", "b"),
			base.NewTransaction("d", "e"),
			base.NewTransaction("a", "b"),
			base.NewTransaction("a", "d", "e"),
			base.NewTransaction("a", "b", "c", "d"),
			base.NewTransaction("a", "b", "c", "d", "e"),
		}, 0.1)
	}
}

func Test_run(t *testing.T) {
	type args struct {
		T []base.Transaction
		s float64
	}
	tests := []struct {
		name string
		args args
		want base.PatternsWithSupport
	}{
		{"test", args{
			[]base.Transaction{
				base.NewTransaction("a", "b", "c"),
				base.NewTransaction("a", "b"),
			},
			0.6,
		}, base.PatternsWithSupport{
			{base.NewPattern("a"), 1},
			{base.NewPattern("b"), 1},
			{base.NewPattern("a", "b"), 1},
		}},
		{"test2", args{
			[]base.Transaction{
				base.NewTransaction("a", "b", "c"),
				base.NewTransaction("a", "b"),
				base.NewTransaction("d", "e"),
				base.NewTransaction("a", "b"),
				base.NewTransaction("a", "d", "e"),
				base.NewTransaction("b", "e"),
			},
			0.5,
		}, base.PatternsWithSupport{
			{base.NewPattern("a"), base.Support(4) / base.Support(6)},
			{base.NewPattern("b"), base.Support(4) / base.Support(6)},
			{base.NewPattern("e"), base.Support(3) / base.Support(6)},
			{base.NewPattern("a", "b"), base.Support(3) / base.Support(6)},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.T, tt.args.s); !got.Equal(tt.want) {
				t.Errorf("Run() = %v (%v), want %v (%v)",
					got.Extract().Set,
					got.ExtractSupport(),
					tt.want.Extract().Set,
					tt.want.ExtractSupport(),
				)
			}
		})
	}
}
