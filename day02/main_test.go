package main

import (
	"bytes"
	"io"
	"testing"
)

func Test_response(t *testing.T) {
	type args struct {
		input io.Reader
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{bytes.NewBufferString("12.00\n20\n8")}, 15},
		{"test1", args{bytes.NewBufferString("15.50\n15\n10")}, 19},
		{"test1", args{bytes.NewBufferString("10.25\n17\n5")}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := response(tt.args.input); got != tt.want {
				t.Errorf("response() = %v, want %v", got, tt.want)
			}
		})
	}
}
