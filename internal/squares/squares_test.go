package squares

import (
	"testing"
)

func TestSquares_binChop(t *testing.T) {
	type args struct {
		value int64
		start int
		end   int
	}
	tests := []struct {
		name  string
		sq    Squares
		args  args
		want  int64
		want1 bool
		want2 int
	}{
		{
			name: "Happy Day",
			sq:   new(5),
			args: args{
				value: 25,
				start: 0,
				end:   6,
			},
			want:  25,
			want1: true,
			want2: 4,
		},
		{
			name: "Unhappy Day",
			sq:   new(5),
			args: args{
				value: 8,
				start: 0,
				end:   6,
			},
			want:  -1,
			want1: false,
			want2: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.sq.binChop(tt.args.value, tt.args.start, tt.args.end)
			if got != tt.want {
				t.Errorf("Squares.binChop() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Squares.binChop() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Squares.binChop() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSquares_FindSumsOfSquares(t *testing.T) {
	type args struct {
		ch chan string
	}
	tests := []struct {
		name string
		sq   Squares
		args args
		want string
		want1 bool
	}{
		{
			name: "Happy day",
			sq: new(4),
			args: args{
				ch: make(chan string, 10),
			},
			want:  "9 (3 * 3) + 16 (4 * 4) = 25 (5 * 5)",
			want1: true,
		},
		{
			name: "Unhappy day",
			sq: new(2),
			args: args{
				ch:   make(chan string, 100),
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sq.FindSumsOfSquares(tt.args.ch)
			got, ok := <-tt.args.ch
			if  ok != tt.want1 {
				t.Errorf("FindSumOfSquares() ok = %v, want %v", ok, tt.want1)
			}
			if got != tt.want {
				t.Errorf("FindSumOfSquares() got = %v, want %v", got, tt.want)
			}
		})
	}
}
