package day3

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_scanCorrupted(t *testing.T) {
	type args struct {
		input            string
		withConditionals bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			},
			want: 161,
		},
		{
			args: args{
				input: common.Input(),
			},
			want: 182619815,
		},
		{
			args: args{
				input:            "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
				withConditionals: true,
			},
			want: 48,
		},
		{
			args: args{
				input:            common.Input(),
				withConditionals: true,
			},
			want: 80747545,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprint(strconv.Itoa(i)), func(t *testing.T) {
			if got := scanCorrupted(tt.args.input, tt.args.withConditionals); got != tt.want {
				t.Errorf("scanCorrupted() = %v, want %v", got, tt.want)
			}
		})
	}
}
