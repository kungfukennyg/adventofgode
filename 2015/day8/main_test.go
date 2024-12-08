package day8

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

const testInput = `""
"abc"
"aaa\"aaa"
"\x27"`

func Test_escape(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 12,
		},
		{
			args: args{common.Input()},
			want: 1369,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := escape(tt.args.input); got != tt.want {
				t.Errorf("escape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unescape(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{testInput},
			want: 19,
		},
		{
			args: args{common.Input()},
			want: 2074,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := unescape(tt.args.input); got != tt.want {
				t.Errorf("unescape() = %v, want %v", got, tt.want)
			}
		})
	}
}
