package apriori

import (
	"reflect"
	"testing"
)

func Benchmark_genL1(b *testing.B) {
	t := largeTransactions()
	tHead, _ := BuildTransactions(t)
	s := 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genL1(*tHead, s)
	}
}

func Benchmark_genC(b *testing.B) {
	l := largePatterns()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genC(l)
	}
}

func Benchmark_genL(b *testing.B) {
	t := largeTransactions()
	tHead, _ := BuildTransactions(t)
	p := largePatterns()
	ps := make([][]string, 0, len(p))
	for k := range p {
		ps = append(ps, split(k))
	}
	s := 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		genL(ps, *tHead, s)
	}
}

func Test_genL(t *testing.T) {
	type args struct {
		c     [][]string
		tHead Transaction
		s     int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "t={{a,b},{a,d},{a,b,c}}, c={{a,b},{a,c},{b,c}} s=2",
			args: args{
				c: [][]string{
					{"a", "b"},
					{"a", "c"},
					{"b", "c"},
				},
				tHead: Transaction{
					items: []string{"a", "b"},
					next: &Transaction{
						items: []string{"a", "d"},
						next: &Transaction{
							items: []string{"a", "b", "c"},
							next:  &Transaction{},
						},
					},
				},
				s: 2,
			},
			want: map[string]int{
				"a,b": 2,
			},
		},
		{
			name: "t={{a,b},{a,d},{a,b,c}}, c={{a,b},{a,c},{b,c}} s=3",
			args: args{
				c: [][]string{
					{"a", "b"},
					{"a", "c"},
					{"b", "c"},
				},
				tHead: Transaction{
					items: []string{"a", "b"},
					next: &Transaction{
						items: []string{"a", "d"},
						next: &Transaction{
							items: []string{"a", "b", "c"},
							next:  &Transaction{},
						},
					},
				},
				s: 3,
			},
			want: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genL(tt.args.c, tt.args.tHead, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genL1(t *testing.T) {
	type args struct {
		tHead Transaction
		s     int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "{{a,b,c},{a,b},{a}} s=2",
			args: args{
				tHead: Transaction{
					items: []string{"a", "b", "c"},
					next: &Transaction{
						items: []string{"a", "b"},
						next: &Transaction{
							items: []string{"a"},
							next:  &Transaction{},
						},
					},
				},
				s: 2,
			},
			want: map[string]int{
				"a": 3,
				"b": 2,
			},
		},
		{
			name: "{{a,b,c},{a,c},{a,d}} s=3",
			args: args{
				tHead: Transaction{
					items: []string{"a", "b", "c"},
					next: &Transaction{
						items: []string{"a", "c"},
						next: &Transaction{
							items: []string{"a", "d"},
							next:  &Transaction{},
						},
					},
				},
				s: 3,
			},
			want: map[string]int{
				"a": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genL1(tt.args.tHead, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genL1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genC(t *testing.T) {
	type args struct {
		l map[string]int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "l={a:3,b:2,c:2}",
			args: args{
				l: map[string]int{
					"a": 3,
					"b": 2,
					"c": 2,
				},
			},
			want: [][]string{
				{"a", "b"},
				{"a", "c"},
				{"b", "c"},
			},
		},
		{
			name: "l={ab:3, ac:2, bc:2, bd:2}",
			args: args{
				l: map[string]int{
					"a,b": 3,
					"a,c": 2,
					"b,c": 2,
					"b,d": 2,
					"c,d": 2,
				},
			},
			want: [][]string{
				{"a", "b", "c"},
				{"b", "c", "d"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genC(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genSubPatterns(t *testing.T) {
	type args struct {
		p []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "p={a,b,c}",
			args: args{
				p: []string{"a", "b", "c"},
			},
			want: [][]string{
				{"b", "c"},
				{"a", "c"},
				{"a", "b"},
			},
		},
		{
			name: "p={a,b,c,d}",
			args: args{
				p: []string{"a", "b", "c", "d"},
			},
			want: [][]string{
				{"b", "c", "d"},
				{"a", "c", "d"},
				{"a", "b", "d"},
				{"a", "b", "c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genSubPatterns(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genSubPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
