package apriori

import (
	"reflect"
	"testing"
)

func BenchmarkMine(b *testing.B) {
	transaction := Transaction{
		items: []string{"a", "b", "c"},
		next: &Transaction{
			items: []string{"a", "b"},
			next: &Transaction{
				items: []string{"a", "c"},
				next: &Transaction{
					items: []string{"a"},
					next:  &Transaction{},
				},
			},
		},
	}
	minSupport := 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mine(transaction, minSupport, false)
	}
}

func TestMine(t *testing.T) {
	type args struct {
		transaction Transaction
		minSupport  int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "{{a,b,c},{a,b},{a,c},{a}}, 2",
			args: args{
				transaction: Transaction{
					items: []string{"a", "b", "c"},
					next: &Transaction{
						items: []string{"a", "b"},
						next: &Transaction{
							items: []string{"a", "c"},
							next: &Transaction{
								items: []string{"a"},
								next:  &Transaction{},
							},
						},
					},
				},
				minSupport: 2,
			},
			want: map[string]int{
				"a":   4,
				"b":   2,
				"c":   2,
				"a,b": 2,
				"a,c": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mine(tt.args.transaction, tt.args.minSupport, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mine() = %v, want %v", got, tt.want)
			}
		})
	}
}
