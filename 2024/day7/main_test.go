package day7

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

const testInput = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_calibrationResult(t *testing.T) {
	type args struct {
		input string
		ops   []op
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: testInput,
				ops:   []op{opAdd, opMult},
			},
			want: 3749,
		},
		{
			args: args{
				input: common.Input(),
				ops:   []op{opAdd, opMult},
			},
			want: 882304362421,
		},
		{
			args: args{
				input: testInput,
				ops:   []op{opAdd, opMult, opConcat},
			},
			want: 11387,
		},
		{
			args: args{
				input: common.Input(),
				ops:   []op{opAdd, opMult, opConcat},
			},
			want: 145149066755184,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := calibrationResult(tt.args.input, tt.args.ops); got != tt.want {
				t.Errorf("calibrationResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
