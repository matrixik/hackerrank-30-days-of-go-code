//

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
		{"", args{bytes.NewBufferString("5")}, 1},
		{"", args{bytes.NewBufferString("13")}, 2},
		{"", args{bytes.NewBufferString("130000000")}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answer(tt.args.input); got != tt.want {
				t.Errorf("answer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAnswer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		answer(bytes.NewBufferString("130000000"))
	}
}
