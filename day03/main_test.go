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
		{"", args{bytes.NewBufferString("3")}, "Weird", false},
		{"", args{bytes.NewBufferString("7")}, "Weird", false},
		{"", args{bytes.NewBufferString("8")}, "Weird", false},
		{"", args{bytes.NewBufferString("4")}, "Not Weird", false},
		{"", args{bytes.NewBufferString("24")}, "Not Weird", false},
		{"", args{bytes.NewBufferString("24\n")}, "Not Weird", false},
		{"", args{bytes.NewBufferString("number")}, "", true},
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
