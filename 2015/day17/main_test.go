package day17

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_uniqueCombos(t *testing.T) {
	type args struct {
		input  string
		liters int
	}
	tests := []struct {
		args args
		want int
		min  int
	}{
		{
			args: args{
				input:  "20\n15\n10\n5\n5",
				liters: 25,
			},
			want: 4,
			min:  3,
		},
		{
			args: args{
				input:  common.Input(),
				liters: 150,
			},
			want: 654,
			min:  57,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, minGot := uniqueCombos(tt.args.input, tt.args.liters); got != tt.want {
				t.Errorf("uniqueCombos() got = %v, want %v", got, tt.want)
			} else if minGot != tt.min {
				t.Errorf("uniqueCombos() min = %v, want %v", minGot, tt.min)

			}
		})
	}
}
