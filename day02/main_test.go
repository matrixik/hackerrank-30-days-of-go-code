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
		{"test2", args{bytes.NewBufferString("15.50\n15\n10")}, 19},
		{"test3", args{bytes.NewBufferString("10.25\n17\n5")}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := response(tt.args.input); got != tt.want {
				t.Errorf("response() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		number float64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"", args{1.2}, 1},
		{"", args{1.5}, 2},
		{"", args{-1.3}, -1},
		{"", args{-1.6}, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.number); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}
