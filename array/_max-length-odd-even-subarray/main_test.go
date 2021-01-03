package main

import "testing"

func TestMaxLenOddEvenSubarray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Maximum length of subarray with alternating elements",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 8, 10}},
			want: 6,
		},
		{
			name: "Maximum length of subarray with alternating elements",
			args: args{arr: []int{10, 12, 14, 7, 8}},
			want: 3,
		},
		{
			name: "Maximum length of subarray with alternating elements",
			args: args{arr: []int{5, 10, 20, 6, 3, 8}},
			want: 3,
		},
		{
			name: "Maximum length of subarray with alternating elements (only odds)",
			args: args{arr: []int{5, 7, 9, 11}},
			want: 1,
		},
		{
			name: "Maximum length of subarray with alternating elements (only even)",
			args: args{arr: []int{2, 4, 6, 8}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxLenOddEvenSubarray(tt.args.arr); got != tt.want {
				t.Errorf("maxLenOddEvenSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
