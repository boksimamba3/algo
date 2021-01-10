package main

import "testing"

func TestMaxSum(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Maximum sum of k consecutive elements",
			args: args{arr: []int{1, 8, 30, -5, 20}, k: 3},
			want: 45,
		},
		{
			name: "Maximum sum of k consecutive elements",
			args: args{arr: []int{5, -10, 6, 90, 3}, k: 2},
			want: 96,
		},
		{
			name: "Maximum sum of k consecutive elements",
			args: args{arr: []int{1, 8, 30, -5, 20}, k: 1},
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
