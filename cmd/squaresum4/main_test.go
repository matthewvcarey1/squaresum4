package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "Happy 1",
			args: []string{"cmd"},
		},
		
		{
			name: "Happy 2",
			args: []string{"cmd","150"},
		},
		
	}
	for _, tt := range tests {
		os.Args = tt.args
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
