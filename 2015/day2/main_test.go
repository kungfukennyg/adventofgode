package day2

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_wrappingPaper(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"2x3x4"},
			want: 58,
		},
		{
			args: args{"1x1x10"},
			want: 43,
		},
		{
			args: args{aoc.Input()},
			want: 1588178,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := wrappingPaper(tt.args.input); got != tt.want {
				t.Errorf("wrappingPaper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ribbonLength(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{"2x3x4"},
			want: 34,
		},
		{
			args: args{"1x1x10"},
			want: 14,
		},
		{
			args: args{aoc.Input()},
			want: 3783758,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := ribbonLength(tt.args.input); got != tt.want {
				t.Errorf("ribbonLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
