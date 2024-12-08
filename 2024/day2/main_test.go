package day2

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_countSafeReports(t *testing.T) {
	type args struct {
		input        string
		withDampener bool
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input:        "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9",
				withDampener: false,
			},
			want: 2,
		},
		{
			args: args{
				input:        common.Input(),
				withDampener: false,
			},
			want: 524,
		},
		{
			args: args{
				input:        "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9",
				withDampener: true,
			},
			want: 4,
		},
		{
			args: args{
				input:        common.Input(),
				withDampener: true,
			},
			want: 569,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := countSafe(tt.args.input, tt.args.withDampener); got != tt.want {
				t.Errorf("countSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}
