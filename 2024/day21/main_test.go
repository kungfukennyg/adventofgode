package day21

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_eval(t *testing.T) {
	type args struct {
		input string
		depth int
	}
	tests := []struct {
		args args
		want int
	}{
// 		{
// 			args: args{
// 				input: `029A
// 980A
// 179A
// 456A
// 379A`,
// 				depth: 3,
// 			},
// 			want: 126384,
// 		},
// 		{
// 			args: args{
// 				input: aoc.Input(),
// 				depth: 3,
// 			},
// 			// too high?
// 			want: 177814,
// 		},
		{
			args: args{
				input: aoc.Input(),
				depth: 27,
			},
			// too high?
			want: 184059,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := eval(tt.args.input, tt.args.depth); got != tt.want {
				t.Errorf("eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
