package day18

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_findExit(t *testing.T) {
	type args struct {
		input  string
		height int
		width  int
		bytes  int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
				height: 7,
				width:  7,
				bytes:  12,
			},
			want: 22,
		},
		{
			args: args{
				input:  aoc.Input(),
				height: 71,
				width:  71,
				bytes:  1024,
			},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := findExit(tt.args.input, tt.args.height, tt.args.width, tt.args.bytes); got != tt.want {
				t.Errorf("findExit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cutsOffExit(t *testing.T) {
	type args struct {
		input  string
		height int
		width  int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				input: `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`,
				height: 7,
				width:  7,
			},
			want: "6,1",
		},
		{
			args: args{
				input:  aoc.Input(),
				height: 71,
				width:  71,
			},
			want: "",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := cutsOffExit(tt.args.input, tt.args.height, tt.args.width); got != tt.want {
				t.Errorf("cutsOffExit() = %v, want %v", got, tt.want)
			}
		})
	}
}
