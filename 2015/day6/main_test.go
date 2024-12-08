package day6

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_lightLevel(t *testing.T) {
	type args struct {
		input  string
		binary bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:  common.Input(),
				binary: true,
			},
			want: 543903,
		},
		{
			args: args{
				input:  common.Input(),
				binary: false,
			},
			want: 14687245,
		},
		{
			args: args{
				input:  "turn on 0,0 through 999,999",
				binary: true,
			},
			want: SIZE * SIZE,
		},
		{
			args: args{
				input:  "turn on 0,0 through 999,999\ntoggle 0,0 through 999,0",
				binary: true,
			},
			want: SIZE * (SIZE - 1),
		},
		{
			args: args{
				input:  "turn on 0,0 through 999,999\ntoggle 0,0 through 999,0\nturn off 499,499 through 500,500",
				binary: true,
			},
			want: SIZE*(SIZE-1) - 4,
		},
		{
			args: args{
				input:  "turn on 0,0 through 0,0",
				binary: false,
			},
			want: 1,
		},
		{
			args: args{
				input:  "toggle 0,0 through 999,999",
				binary: false,
			},
			want: 2 * SIZE * SIZE,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := lightLevel(tt.args.input, tt.args.binary); got != tt.want {
				t.Errorf("lightLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
