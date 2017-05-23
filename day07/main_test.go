// https://www.hackerrank.com/challenges/30-array

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
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"", args{bytes.NewBufferString("4\n1 4 3 2")}, "2 3 4 1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := answer(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("answer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("answer() = %v, want %v", got, tt.want)
			}
		})
	}
}
