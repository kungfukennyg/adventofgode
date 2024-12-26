package day14

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_safetyFactor(t *testing.T) {
	type args struct {
		input   string
		height  int
		width   int
		seconds int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
				height:  7,
				width:   11,
				seconds: 100,
			},
			want: 12,
		},
		{
			args: args{
				input:   aoc.Input(),
				height:  103,
				width:   101,
				seconds: 100,
			},
			want: 220971520,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := safetyFactor(tt.args.input, tt.args.height, tt.args.width, tt.args.seconds); got != tt.want {
				t.Errorf("safetyFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_christmasTree(t *testing.T) {
	type args struct {
		input  string
		height int
		width  int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:  aoc.Input(),
				height: 103,
				width:  101,
			},
			want: 220971520,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := christmasTree(tt.args.input, tt.args.height, tt.args.width); got != tt.want {
				t.Errorf("christmasTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
