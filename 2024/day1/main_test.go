package day1

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_totalDistance(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"},
			want: 11,
		},
		{
			args: args{aoc.Input()},
			want: 2192892,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := totalDistance(tt.args.input); got != tt.want {
				t.Errorf("totalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarityScore(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: "3   4\n4   3\n2   5\n1   3\n3   9\n3   3"},
			want: 31,
		},
		{
			args: args{aoc.Input()},
			want: 22962826,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := similarityScore(tt.args.input); got != tt.want {
				t.Errorf("similarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
