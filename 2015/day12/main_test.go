package day12

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_sumNumbers(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: "[1,2,3]"},
			want: 6,
		},
		{
			args: args{input: `{"a":2,"b":4}"}`},
			want: 6,
		},
		{
			args: args{input: "[[[[3]]]"},
			want: 3,
		},
		{
			args: args{input: `{"a":{"b":4},"c":-1}`},
			want: 3,
		},
		{
			args: args{input: `[]`},
			want: 0,
		},
		{
			args: args{input: `{}`},
			want: 0,
		},
		{
			args: args{input: aoc.Input()},
			want: 156366,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sumNumbers(tt.args.input); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumSkipRed(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{`{"a": [1,2,3]}`},
			want: 6,
		},
		{
			args: args{`[1,{"c":"red","b":2},3]`},
			want: 4,
		},
		{
			args: args{`{"d":"red","e":[1,2,3,4],"f":5}`},
			want: 0,
		},
		{
			args: args{`[1,"red",5]`},
			want: 6,
		},
		{
			args: args{aoc.Input()},
			want: 96852,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sumSkipRed(tt.args.input); got != tt.want {
				t.Errorf("sumSkipRed() = %v, want %v", got, tt.want)
			}
		})
	}
}
