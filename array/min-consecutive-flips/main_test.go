package main

import "testing"

func TestMinConsecutiveFlipsCount(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Minimum consecutive flips",
			args: args{arr: []int{1, 0, 0, 0, 1, 0, 1, 0, 0, 0}},
			want: 3,
		},
		{
			name: "Minimum consecutive flips",
			args: args{arr: []int{1, 1, 0, 0, 0, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minConsecutiveFlipsCount(tt.args.arr); got != tt.want {
				t.Errorf("minConsecutiveFlipsCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
