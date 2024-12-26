package day13

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_minTokens(t *testing.T) {
	type args struct {
		input    string
		prizeMod int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`},
			want: 480,
		},
		{
			args: args{input: aoc.Input()},
			want: 30413,
		},
		{
			args: args{input: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`, prizeMod: 10000000000000},
			want: 875318608908,
		},
		{
			args: args{input: aoc.Input(), prizeMod: 10000000000000},
			want: 92827349540204,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := minTokens(tt.args.input, tt.args.prizeMod); got != tt.want {
				t.Errorf("minTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}
