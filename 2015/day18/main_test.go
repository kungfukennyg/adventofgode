package day18

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_sim(t *testing.T) {
	type args struct {
		input           string
		steps           int
		cornersAlwaysOn bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
				steps: 4,
			},
			want: 4,
		},
		{
			args: args{
				input: common.Input(),
				steps: 100,
			},
			want: 821,
		},
		{
			args: args{
				input: `.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
				steps:           5,
				cornersAlwaysOn: true,
			},
			want: 17,
		},
		{
			args: args{
				input:           common.Input(),
				steps:           100,
				cornersAlwaysOn: true,
			},
			want: 886,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sim(tt.args.input, tt.args.steps, tt.args.cornersAlwaysOn); got != tt.want {
				t.Errorf("sim() = %v, want %v", got, tt.want)
			}
		})
	}
}
