package day6

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

const testInput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func Test_visitedByGuard(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 41,
		},
		{
			args: args{aoc.Input()},
			want: 4580,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := parse(tt.args.input)
			if got := b.visitedByGuard(); got != tt.want {
				t.Errorf("visitedByGuard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_infiniteObstacles(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 6,
		},
		{
			args: args{aoc.Input()},
			want: 1480,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			b := parse(tt.args.input)
			if got := b.infiniteObstacles(); got != tt.want {
				t.Errorf("infiniteObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
