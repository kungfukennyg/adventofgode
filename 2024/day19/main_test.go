package day19

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_matchPatterns(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args     args
		possible int
		ways     int
	}{
		{
			args: args{`r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`},
			possible: 6,
			ways:     16,
		},
		{
			args:     args{aoc.Input()},
			possible: 209,
			ways:     777669668613191,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotPossible, gotWays := matchPatterns(tt.args.input)
			if gotPossible != tt.possible {
				t.Errorf("matchPatterns(), possible = %v, want %v", gotPossible, tt.possible)
			}
			if gotWays != tt.ways {
				t.Errorf("matchPatterns, ways = %v, want %v", gotWays, tt.ways)
			}
		})
	}
}
