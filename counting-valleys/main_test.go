package main

import "testing"

func Test_countingValleys(t *testing.T) {
	type args struct {
		n int32
		s string
	}
	tests := []struct {
		name string
		args args
		want int32
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countingValleys(tt.args.n, tt.args.s); got != tt.want {
				t.Errorf("countingValleys() = %v, want %v", got, tt.want)
			}
		})
	}
}