package main

import "testing"

func TestGetSum(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Get sum",
			args: args{
				start: 0,
				end:   2,
			},
			want: 13,
		},
		{
			name: "Get sum",
			args: args{
				start: 1,
				end:   3,
			},
			want: 20,
		},
		{
			name: "Get sum",
			args: args{
				start: 2,
				end:   6,
			},
			want: 27,
		},
	}
	arr := []int{2, 8, 3, 9, 6, 5, 4}
	getSum := prefixSum(arr)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
