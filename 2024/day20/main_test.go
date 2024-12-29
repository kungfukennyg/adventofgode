package day20

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_cheats(t *testing.T) {
	type args struct {
		input        string
		cheatDepth   int
		desiredDelta int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
				cheatDepth:   2,
				desiredDelta: 1,
			},
			want: 44,
		},
		{
			args: args{
				input: `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`,
				cheatDepth:   20,
				desiredDelta: 50,
			},
			want: 285,
		},
		{
			args: args{
				input:        aoc.Input(),
				cheatDepth:   2,
				desiredDelta: 100,
			},
			want: 1289,
		},
		{
			args: args{
				input:        aoc.Input(),
				cheatDepth:   20,
				desiredDelta: 100,
			},
			want: 982425,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := cheats(tt.args.input, tt.args.cheatDepth, tt.args.desiredDelta); got != tt.want {
				t.Errorf("cheats() = %v, want %v", got, tt.want)
			}
		})
	}
}
