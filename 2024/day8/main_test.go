package day8

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

const testInput = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func Test_findAntinodes(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 14,
		},
		{
			args: args{aoc.Input()},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findAntinodes(tt.args.input); got != tt.want {
				t.Errorf("findAntinodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

const part2Input = `T.........
...T......
.T........
..........
..........
..........
..........
..........
..........
..........`

func Test_findAntinodesAnyDistance(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 34,
		},
		{
			args: args{part2Input},
			want: 9,
		},
		{
			args: args{aoc.Input()},
			want: 809,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findAntinodesAnyDistance(tt.args.input); got != tt.want {
				t.Errorf("findAntinodesAnyDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
