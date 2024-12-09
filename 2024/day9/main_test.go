package day9

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_processDisk(t *testing.T) {
	type args struct {
		input     string
		wholeFile bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{input: "12345", wholeFile: false},
			want: 60,
		},
		{
			args: args{input: "2333133121414131402", wholeFile: false},
			want: 1928,
		},
		{
			args: args{input: common.Input(), wholeFile: false},
			want: 6353658451014,
		},
		{
			args: args{input: "2333133121414131402", wholeFile: true},
			want: 2858,
		},
		{
			args: args{input: common.Input(), wholeFile: true},
			want: 6382582136592,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := processDisk(tt.args.input, tt.args.wholeFile); got != tt.want {
				t.Errorf("processDisk() = %v, want %v", got, tt.want)
			}
		})
	}
}
