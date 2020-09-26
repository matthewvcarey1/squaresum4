package main

import (
	"testing"
)

func Test_binchop(t *testing.T) {
	type args struct {
		value int64
		set   []int64
		index int
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
		want2 int
	}{
		{
			name: "Happy day",
			args: args{
				value: 5,
				set:   []int64{2, 3, 4, 5, 6},
				index: 0,
			},
			want:  5,
			want1: true,
			want2: 3,
		},
		{
			name: "Unhappy day 1",
			args: args{
				value: 5,
				set:   []int64{2, 3, 4, 6, 7, 12},
				index: 0,
			},
			want:  -1,
			want1: false,
			want2: -1,
		},
		{
			name: "Unhappy day 2",
			args: args{
				value: 1,
				set:   []int64{2, 3, 4, 6, 7, 12},
				index: 0,
			},
			want:  -1,
			want1: false,
			want2: -1,
		},
		{
			name: "Unhappy day 3",
			args: args{
				value: 200,
				set:   []int64{2, 3, 4, 6, 7, 12},
				index: 0,
			},
			want:  -1,
			want1: false,
			want2: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := binchop(tt.args.value, tt.args.set, tt.args.index)
			if got != tt.want {
				t.Errorf("binchop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("binchop() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("binchop() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_findSumOfSquares(t *testing.T) {
	type args struct {
		size int
		ch   chan string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Happy day",
			args: args{
				size: 4,
				ch: make(chan string),
			},
			want:"9 (3 * 3) + 16 (4 * 4) = 25 (5 * 5)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go findSumOfSquares(tt.args.size, tt.args.ch)
			got := <- tt.args.ch
			if got != tt.want {
				t.Errorf("findSumOfSquares() got = %v, want %v", got, tt.want)
			}
		})
	}
}
