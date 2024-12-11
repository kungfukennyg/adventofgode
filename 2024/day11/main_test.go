package day11

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_sim(t *testing.T) {
	type args struct {
		input string
		steps int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: "125 17",
				steps: 25,
			},
			want: 55312,
		},
		{
			args: args{
				input: common.Input(),
				steps: 75,
			},
			want: 221632504974231,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sim(tt.args.input, tt.args.steps); got != tt.want {
				t.Errorf("sim() = %v, want %v", got, tt.want)
			}
		})
	}
}
