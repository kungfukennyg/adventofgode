package day12

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_fencePrice(t *testing.T) {
	type args struct {
		input    string
		discount bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: `AAAA
BBCD
BBCC
EEEC`},
			want: 140,
		},
		{
			args: args{input: `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`},
			want: 772,
		},
		{
			args: args{input: `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`},
			want: 1930,
		},
		{
			args: args{input: aoc.Input()},
			want: 1363484,
		},
		{
			args: args{input: `AAAA
BBCD
BBCC
EEEC`, discount: true},
			want: 80,
		},
		{
			args: args{input: `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`, discount: true},
			want: 236,
		},
		{
			args: args{input: `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`, discount: true},
			want: 368,
		},
		{
			args: args{input: aoc.Input(), discount: true},
			want: 838988,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := fencePrice(tt.args.input, tt.args.discount); got != tt.want {
				t.Errorf("fencePrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
