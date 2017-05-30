// https://www.hackerrank.com/challenges/30-2d-arrays

package main

import (
	"bytes"
	"io"
	"testing"
)

func Test_answer(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{bytes.NewBufferString(`1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0`)}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answer(tt.args.input); got != tt.want {
				t.Errorf("answer() = %v, want %v", got, tt.want)
			}
		})
	}
}
