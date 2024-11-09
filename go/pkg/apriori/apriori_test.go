package apriori

import (
	"data-mining/pkg/base"
	"math/rand"
	"testing"
)

func BenchmarkRun(b *testing.B) {
	T := generateLargeTransactions(10000) // Generate 50,000 transactions
	s := base.Support(0.02)               // Use a tiny minSupport
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Run(T, s)
	}
}

func generateLargeTransactions(n int) []base.Transaction {
	var T []base.Transaction
	for i := 0; i < n; i++ {
		items := make([]string, 20)
		for j := 0; j < 20; j++ {
			items[j] = RandomString(5)
		}
		T = append(T, base.NewTransaction(items...))
	}
	return T
}

func RandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func Test_run(t *testing.T) {
	type args struct {
		T []base.Transaction
		s base.Support
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
			base.NewPattern("a"):      1,
			base.NewPattern("b"):      1,
			base.NewPattern("a", "b"): 1,
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
			base.NewPattern("a"):      base.Support(4) / base.Support(6),
			base.NewPattern("b"):      base.Support(4) / base.Support(6),
			base.NewPattern("e"):      base.Support(3) / base.Support(6),
			base.NewPattern("a", "b"): base.Support(3) / base.Support(6),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.T, tt.args.s); !got.Equal(tt.want) {
				t.Errorf("Run() = %v, want %v",
					got,
					tt.want,
				)
			}
		})
	}
}
