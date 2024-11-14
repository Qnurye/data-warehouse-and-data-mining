package apriori

import (
	"reflect"
	"testing"
)

func Benchmark_contains(b *testing.B) {
	transaction := randomStrings(1000)
	pattern := randomStrings(10)
	for i := 0; i < b.N; i++ {
		contains(transaction, pattern)
	}
}

func Benchmark_canMerge(b *testing.B) {
	p1 := randomStrings(10)
	p2 := randomStrings(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		canMerge(p1, p2)
	}
}

func Benchmark_merge(b *testing.B) {
	x := randomMap(100)
	y := randomMap(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merge(x, y)
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

func Test_canMerge(t *testing.T) {
	type args struct {
		p1 []string
		p2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "p1={a,b} p2={a,c}",
			args: args{
				p1: []string{"a", "b"},
				p2: []string{"a", "c"},
			},
			want: true,
		},
		{
			name: "p1={a,b} p2={b,c}",
			args: args{
				p1: []string{"a", "b"},
				p2: []string{"b", "c"},
			},
			want: false,
		},
		{
			name: "p1={a,b,c} p2={a,b,d}",
			args: args{
				p1: []string{"a", "b", "c"},
				p2: []string{"a", "b", "d"},
			},
			want: true,
		},
		{
			name: "p1={a,b,c} p2={a,c,d}",
			args: args{
				p1: []string{"a", "b", "c"},
				p2: []string{"a", "c", "d"},
			},
			want: false,
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

func Test_merge(t *testing.T) {
	type args struct {
		a map[string]int
		b map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"{} + {} = {}",
			args{map[string]int{}, map[string]int{}},
			map[string]int{},
		},
		{
			"{'a': 1} + {} = {'a': 1}",
			args{map[string]int{"a": 1}, map[string]int{}},
			map[string]int{"a": 1},
		},
		{
			"{} + {'a': 1} = {'a': 1}",
			args{map[string]int{}, map[string]int{"a": 1}},
			map[string]int{"a": 1},
		},
		{
			"{'a': 1} + {'b': 2} = {'a': 1, 'b': 2}",
			args{map[string]int{"a": 1}, map[string]int{"b": 2}},
			map[string]int{"a": 1, "b": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
