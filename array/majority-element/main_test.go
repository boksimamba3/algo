package main

import "testing"

func TestFindMajority(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Majority element",
			args: args{arr: []int{8, 3, 4, 8, 8}},
			want: 3,
		},
		{
			name: "Majority element (no majority)",
			args: args{arr: []int{3, 7, 4, 7, 7, 5}},
			want: -1,
		},
		{
			name: "Majority element",
			args: args{arr: []int{3, 7, 4, 7, 7, 5, 7}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMajority(tt.args.arr); got != tt.want {
				t.Errorf("findMajority() = %v, want %v", got, tt.want)
			}
		})
	}
}
