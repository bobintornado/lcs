package lcs

import "testing"

type args struct {
	originResult []string
	newResult    []string
}

type testCase struct {
	name string
	args args
	want float64
}

var tests []testCase = []testCase{
	{
		name: "test 1",
		args: args{
			originResult: []string{"L1", "L2", "L3", "L4", "L5"},
			newResult:    []string{"L2", "L3", "L4", "L5"},
		},
		want: 0.8,
	},
	{
		name: "test 2",
		args: args{
			originResult: []string{"L1", "L2", "L3", "L4", "L5"},
			newResult:    []string{"L4", "L3", "L2", "L5"},
		},
		want: 0.4,
	},
	{
		name: "test 3",
		args: args{
			originResult: []string{"L1", "L2", "L3", "L4", "L5"},
			newResult:    []string{"L5", "L4", "L3", "L2", "L1"},
		},
		want: 0.2,
	},
	{
		name: "test 4",
		args: args{
			originResult: []string{"L1", "L2", "L3", "L4", "L5"},
			newResult:    []string{"L2", "L3", "L12", "L13", "L4", "L5"},
		},
		want: 0.8,
	},
	{
		name: "test 5",
		args: args{
			originResult: []string{"L1", "L2", "L3", "L4", "L5"},
			newResult:    []string{"L12", "L13", "L14", "L15"},
		},
		want: 0.0,
	},
	{
		name: "test 6",
		args: args{
			originResult: []string{},
			newResult:    []string{"L12", "L13", "L14", "L15"},
		},
		want: 0.0,
	},
	{
		name: "test 7",
		args: args{
			originResult: []string{"L1", "L2", "L9", "L4", "L5"},
			newResult:    []string{"L1", "L2", "L3", "L4", "L5"},
		},
		want: 0.8,
	},
}

func TestComparator_Evaluate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Comparator{}
			if got := r.Evaluate(tt.args.originResult, tt.args.newResult); got != tt.want {
				t.Errorf("Comparator.Evaluate() = %v, want %v", got, tt.want)
			}
		})
	}
}

var result float64

func BenchmarkComparator_Evaluate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var ans float64
		for _, tt := range tests {
			r := Comparator{}
			ans = r.Evaluate(tt.args.originResult, tt.args.newResult)
		}
		result = ans
	}
}

// func TestMaxInt(t *testing.T) {
// 	type args struct {
// 		a int
// 		b int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestLcsLength(t *testing.T) {
// 	type args struct {
// 		a []string
// 		b []string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := LcsLength(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("LcsLength() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
