package day15

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_bestRatio(t *testing.T) {
	type args struct {
		input     string
		teaspoons int
		calories  int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
				teaspoons: 100,
				calories:  -1,
			},
			want: 62842880,
		},
		{
			args: args{
				input:     aoc.Input(),
				teaspoons: 100,
				calories:  -1,
			},
			want: 21367368,
		},
		{
			args: args{
				input: `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
		Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`,
				teaspoons: 100,
				calories:  500,
			},
			want: 57600000,
		},
		{
			args: args{
				input:     aoc.Input(),
				teaspoons: 100,
				calories:  500,
			},
			want: 21367368,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := bestRecipe(tt.args.input, tt.args.teaspoons, tt.args.calories); got != tt.want {
				t.Errorf("bestRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}
