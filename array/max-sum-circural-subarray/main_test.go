package main

import "testing"

func TestMaxSumCircuralSubarray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Maximum sum of circural subarray",
			args: args{arr: []int{5, -2, 3, 4}},
			want: 12,
		},
		{
			name: "Maximum sum of circural subarray",
			args: args{arr: []int{8, -4, 3, -5, 4}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSumCircuralSubarray(tt.args.arr); got != tt.want {
				t.Errorf("maxSumCircuralSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
