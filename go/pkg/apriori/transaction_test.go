package apriori

import (
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkBuildTransactions(b *testing.B) {
	transactions := largeTransactions()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BuildTransactions(transactions)
	}
}

func Benchmark_sortItems(b *testing.B) {
	items := randomStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sortItems(items)
	}
}

func TestBuildTransactions(t *testing.T) {
	type args struct {
		transactions [][]string
	}
	tests := []struct {
		name  string
		args  args
		want  *Transaction
		want1 int
	}{
		{
			"{{a, b, c}, {a, b}, {a, b, c}} => {a, b, c} -> {a, b} -> {a, b, c}",
			args{[][]string{{"a", "b", "c"}, {"a", "b"}, {"a", "b", "c"}}},
			&Transaction{
				items: []string{"a", "b", "c"},
				next: &Transaction{
					items: []string{"a", "b"},
					next: &Transaction{
						items: []string{"a", "b", "c"},
						next:  &Transaction{},
					},
				},
			},
			3,
		},
		{
			"{{a, b, c}, {a, b}, {b, c, a}} => {a, b, c} -> {a, b} -> {a, b, c}",
			args{[][]string{{"a", "b", "c"}, {"a", "b"}, {"b", "c", "a"}}},
			&Transaction{
				items: []string{"a", "b", "c"},
				next: &Transaction{
					items: []string{"a", "b"},
					next: &Transaction{
						items: []string{"a", "b", "c"},
						next:  &Transaction{},
					},
				},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := BuildTransactions(tt.args.transactions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTransactions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("BuildTransactions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_sortItems(t *testing.T) {
	type args struct {
		items []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"{a, 1, c, string} => {1, a, c, string}",
			args{[]string{"a", "1", "c", "string"}},
			[]string{"1", "a", "c", "string"},
		},
		{
			"{1, 2, 4, 3} => {1, 2, 3, 4}",
			args{[]string{"1", "2", "4", "3"}},
			[]string{"1", "2", "3", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func randomStrings(l int) []string {
	strDict := "abcdef"
	var s []string
	for i := 0; i < l; i++ {
		s = append(s, string(strDict[rand.Int()%len(strDict)]))
	}
	return s
}

func largeTransactions() [][]string {
	var t [][]string
	for i := 0; i < 1000; i++ {
		t = append(t, randomStrings(3))
	}
	return t
}

func largePatterns() map[string]int {
	p := make(map[string]int)
	for i := 0; i < 50; i++ {
		p[join(randomStrings(3))] = rand.Int() % 100
	}
	return p
}

func randomMap(l int) map[string]int {
	m := make(map[string]int)
	for i := 0; i < l; i++ {
		m[join(randomStrings(3))] = rand.Int() % 100
	}
	return m
}
