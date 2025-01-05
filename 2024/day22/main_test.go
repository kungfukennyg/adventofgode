package day22

import (
	"strconv"
	"testing"

	"github.com/kungfukennyg/adventofgode/aoc"
)

func Test_sumSecrets(t *testing.T) {
	type args struct {
		input string
		depth int
	}
	tests := []struct {
		args args
		want int64
	}{
		{
			args: args{
				input: "123",
				depth: 10,
			},
			want: 5908254,
		},
		{
			args: args{
				input: `1
10
100
2024`,
				depth: 2000,
			},
			want: 37327623,
		},
		{
			args: args{
				input: aoc.Input(),
				depth: 2000,
			},
			want: 13584398738,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := sumSecrets(tt.args.input, tt.args.depth); got != tt.want {
				t.Errorf("sumSecrets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mostBananas(t *testing.T) {
	type args struct {
		input string
		depth int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				input: `1
2
3
2024`,
				depth: 2000,
			},
			want: 23,
		},
		{
			args: args{
				input: aoc.Input(),
				depth: 2000,
			},
			want: 1612,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := mostBananas(tt.args.input, tt.args.depth); got != tt.want {
				t.Errorf("mostBananas() = %v, want %v", got, tt.want)
			}
		})
	}
}