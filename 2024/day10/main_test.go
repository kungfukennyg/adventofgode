package day10

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_scoreTrailheads(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args   args
		unique int
		total  int
	}{
		{
			args: args{input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`},
			unique: 36,
			total:  81,
		},
		{
			args:   args{aoc.Input()},
			unique: 624,
			total:  1483,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			unique, total := scoreTrailheads(tt.args.input)
			if unique != tt.unique {
				t.Errorf("scoreTrailheads() unique = %v, want %v", unique, tt.unique)
			}
			if total != tt.total {
				t.Errorf("scoreTrailheads() total = %v, want %v", total, tt.total)
			}
		})
	}
}
