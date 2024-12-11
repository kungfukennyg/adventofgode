package day16

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/common"
)

func Test_whichSue(t *testing.T) {
	type args struct {
		input           string
		compounds       string
		greater, lesser []string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: common.Input(),
				compounds: `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`,
			},
			want: 40,
		},
		{
			args: args{
				input:   common.Input(),
				greater: []string{"cats", "trees"},
				lesser:  []string{"pomeranians", "goldfish"},
				compounds: `children: 3
cats: 7
samoyeds: 2
pomeranians: 3
akitas: 0
vizslas: 0
goldfish: 5
trees: 3
cars: 2
perfumes: 1`,
			},
			want: 40,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := whichSue(tt.args.input, tt.args.compounds, tt.args.greater, tt.args.lesser); got != tt.want {
				t.Errorf("whichSue() = %v, want %v", got, tt.want)
			}
		})
	}
}
