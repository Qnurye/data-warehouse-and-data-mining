package apriori

import "testing"

func Benchmark_contains(b *testing.B) {
	transaction := randomStrings(1000)
	pattern := randomStrings(10)
	for i := 0; i < b.N; i++ {
		contains(transaction, pattern)
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		transaction []string
		pattern     []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"{'a', 'b', 'c'} contains {'a', 'b'}",
			args{[]string{"a", "b", "c"}, []string{"a", "b"}},
			true,
		},
		{
			"{'a', 'b'} contains {'a', 'b'}",
			args{[]string{"a", "b"}, []string{"a", "b"}},
			true,
		},
		{
			"{'a', 'b', 'c'} contains {'a', 'c'}",
			args{[]string{"a", "b", "c"}, []string{"a", "c"}},
			true,
		},
		{
			"{'a', 'b', 'c'} does not contain {'d'}",
			args{[]string{"a", "b", "c"}, []string{"d"}},
			false,
		},
		{
			"{'a', 'b', 'c'} does not contain {'a', 'b', 'c', 'd'}",
			args{[]string{"a", "b", "c"}, []string{"a", "b", "c", "d"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.transaction, tt.args.pattern); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
