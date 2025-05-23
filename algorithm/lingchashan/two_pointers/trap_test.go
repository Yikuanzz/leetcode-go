package twopointers

import (
	"testing"
)

func Test_trap_1(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap_1(tt.args.height); got != tt.want {
				t.Errorf("trap_1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trap_2(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap_2(tt.args.height); got != tt.want {
				t.Errorf("trap_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
