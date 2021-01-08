package main

import "testing"

func TestHasEquilibriumPoint(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Has equilibrium point",
			args: args{arr: []int{3, 4, 8, -9, 20, 6}},
			want: true,
		},
		{
			name: "Has equilibrium point",
			args: args{arr: []int{-7, 1, 5, 2, -4, 3, 0}},
			want: true,
		},
		{
			name: "Doesn't have equilibrium point",
			args: args{arr: []int{2, 3, 5, 7}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasEquilibriumPoint(tt.args.arr); got != tt.want {
				t.Errorf("hasEquilibriumPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
